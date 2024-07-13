package controller

import (
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

func handleEntry(g *echo.Group) {
    r := g.Group("/entry")
    r.POST("", getEntries)
}

