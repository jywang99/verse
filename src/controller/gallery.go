package controller

import (
	"github.com/labstack/echo/v4"
	"jy.org/verse/src/service"
)

func getThumb(c echo.Context) error {
    tPath := c.QueryParam("path")

    file, mime, err := service.GetThumb(tPath)
    if err != nil {
        return handleError(c, err)
    }

    return c.Stream(200, mime, file)
}

func getCastPic(c echo.Context) error {
    tPath := c.QueryParam("path")

    file, mime, err := service.GetCastPic(tPath)
    if err != nil {
        return handleError(c, err)
    }

    return c.Stream(200, mime, file)
}

func handleThumb(g *echo.Group) {
    r := g.Group("/gallery")
    r.GET("/thumb", getThumb)
    r.GET("/cast", getCastPic)
}

