package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"jy.org/verse/src/entity"
	"jy.org/verse/src/service"
)

func getEntries(c echo.Context) error {
    query := entity.DefaultGetEntries()
    if err := c.Bind(query); err != nil {
        return err
    }

    if err := query.Validate(); err != nil {
        return handleError(c, err)
    }

    entries, err := service.GetEntries(*query)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, entries)
}

func getEntry(c echo.Context) error {
    id, err := parseIdParam(c, "id")
    if err != nil {
        return handleError(c, err)
    }

    entry, err := service.GetEntry(id)
    if err != nil {
        return handleError(c, err)
    }

    token, err := genMediaToken(entry.Path)
    if err != nil {
        return handleError(c, err)
    }

    cookie := http.Cookie{
        Name: fmt.Sprintf("stream-%d", id),
        Value: token,
	SameSite: http.SameSiteStrictMode,
        Expires: time.Now().Add(time.Second * authCfg.MediaTokenDuration),
        Path: fmt.Sprintf("/media/stream/%d", id),
    }
    c.SetCookie(&cookie)
    cookie = http.Cookie{
        Name: fmt.Sprintf("static-%d", id),
        Value: token,
	SameSite: http.SameSiteStrictMode,
        Expires: time.Now().Add(time.Second * authCfg.MediaTokenDuration),
        Path: fmt.Sprintf("/media/static/%d", id),
    }
    c.SetCookie(&cookie)

    return c.JSON(200, entry)
}

func handleEntry(g *echo.Group) {
    r := g.Group("/entry")
    r.POST("", getEntries)
    r.GET("/:id", getEntry)
}

