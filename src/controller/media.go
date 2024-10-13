package controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"jy.org/verse/src/config"
	"jy.org/verse/src/except"
	"jy.org/verse/src/service"
)

var mediaSignKey = []byte(config.Config.Auth.MediaSecret)

type jwtMediaClaims struct {
    jwt.RegisteredClaims
    AllowedPath string `json:"allowedPath"`
}

func genMediaToken(path string) (string, error) {
    claims := &jwtMediaClaims{
        AllowedPath: path,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    ts, err := token.SignedString(mediaSignKey)
    if err != nil {
        logger.ERROR.Println("Failed to sign token: ", err)
        return ts, errors.New("Failed to sign token")
    }

    return ts, nil
}

func getTokenFromCookie(c echo.Context, cname string) (*jwtMediaClaims, error) {
    cookie, err := c.Cookie(cname)
    if err != nil {
        return nil, err
    }

    token := cookie.Value
    claims := &jwtMediaClaims{}
    _, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return mediaSignKey, nil
    })

    return claims, err
}

// TODO refactor
func serveVideo(c echo.Context) error {
    id, err := parseIdParam(c, "id")
    if err != nil {
        return handleError(c, err)
    }

    claims, err := getTokenFromCookie(c, fmt.Sprintf("stream-%d", id))
    if err != nil {
        return handleError(c, except.NewHandledError(except.AuthErr, "Failed to get token"))
    }

    // get path from token
    path := claims.AllowedPath
    subPath := c.QueryParam("file")
    logger.INFO.Println("Serving video: ", path, subPath)

    // Open the video file
    file, err := os.Open(filepath.Join(config.Config.File.MediaRoot, path, subPath))
    if err != nil {
        return c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening video file: %v", err))
    }
    defer file.Close()

    // Get the file information
    fileInfo, err := file.Stat()
    if err != nil {
        return c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file information: %v", err))
    }

    // Set the content type
    c.Response().Header().Set("Content-Type", "video/mp4")

    // Get the Range header from the request
    rangeHeader := c.Request().Header.Get("Range")
    if rangeHeader == "" {
        return handleError(c, except.NewHandledError(except.BadRequestErr, "Range header not found"))
    }

    // Parse the Range header
    rangeParts := strings.Split(strings.TrimPrefix(rangeHeader, "bytes="), "-")
    if len(rangeParts) != 2 {
        return handleError(c, except.NewHandledError(except.BadRequestErr, "Invalid range header"))
    }
    start, err := strconv.ParseInt(rangeParts[0], 10, 64)
    if err != nil {
        return c.String(http.StatusRequestedRangeNotSatisfiable, "Invalid range start")
    }
    end := fileInfo.Size() - 1
    if rangeParts[1] != "" {
        end, err = strconv.ParseInt(rangeParts[1], 10, 64)
        if err != nil {
            return c.String(http.StatusRequestedRangeNotSatisfiable, "Invalid range end")
        }
    }
    if start > end || end >= fileInfo.Size() {
        return c.String(http.StatusRequestedRangeNotSatisfiable, "Range not satisfiable")
    }

    // Set the content range header
    c.Response().Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size()))
    c.Response().Header().Set("Accept-Ranges", "bytes")
    c.Response().WriteHeader(http.StatusPartialContent)

    // Seek to the start position and stream the requested range
    // service.StreamVideo(path, subPath, end, c.Response().Writer)
    _, err = file.Seek(start, 0)
    if err != nil {
        return c.String(http.StatusInternalServerError, fmt.Sprintf("Error seeking file: %v", err))
    }
    buffer := make([]byte, 1024*1024) // 1MB buffer size
    for {
        n, err := file.Read(buffer)
        if err != nil && err.Error() != "EOF" {
            return c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading file: %v", err))
        }
        if n == 0 {
            break
        }
        if _, err := c.Response().Writer.Write(buffer[:n]); err != nil {
            return c.String(http.StatusInternalServerError, fmt.Sprintf("Error writing file: %v", err))
        }
        c.Response().Flush()
    }

    return nil
}

func getFullMedia(c echo.Context) error {
    id, err := parseIdParam(c, "id")
    if err != nil {
	return handleError(c, err)
    }

    claims, err := getTokenFromCookie(c, fmt.Sprintf("static-%d", id))
    if err != nil {
        return handleError(c, except.NewHandledError(except.AuthErr, "Failed to get token"))
    }
    // get path from token
	path := claims.AllowedPath
    file := c.QueryParam("file")

    content, ftype, err := service.GetStaticContent(path, file)
    if err != nil {
        return handleError(c, err)
    }

    return c.Stream(http.StatusOK, ftype, content)
}

func handleMedia(e *echo.Echo) { // TODO unique url for each entry
    r := e.Group("/media")
    r.GET("/static/:id", getFullMedia)
    r.GET("/stream/:id", serveVideo)
}

