package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"jy.org/verse/src/config"
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

func getPartialContent(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtMediaClaims)
	path := claims.AllowedPath
    file := c.QueryParam("file")

    err := service.SeekVideo(path, file)
    if err != nil {
        return handleError(c, err)
    }

    return c.String(http.StatusOK, file) // TODO
}

func getStaticContent(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtMediaClaims)
	path := claims.AllowedPath
    file := c.QueryParam("file")

    content, ftype, err := service.GetStaticContent(path, file)
    if err != nil {
        return handleError(c, err)
    }

    return c.Stream(http.StatusOK, ftype, content)
}

func handleMedia(e *echo.Echo) {
    r := e.Group("/media")
    config := echojwt.Config{
        NewClaimsFunc: func(c echo.Context) jwt.Claims {
            return new(jwtMediaClaims)
        },
        SigningKey: mediaSignKey,
    }
    r.Use(echojwt.WithConfig(config))
    r.GET("/partial", getPartialContent)
    r.GET("/static", getStaticContent)
}

