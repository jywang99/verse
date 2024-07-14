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

func getEntry(c echo.Context) error {
    id, err := parseIdParam(c, "id")
    if err != nil {
        return handleError(c, err)
    }

    entry, err := service.GetEntry(id)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, entry)
}

func handleEntry(g *echo.Group) {
    r := g.Group("/entry")
    r.POST("", getEntries)
    r.GET("/:id", getEntry)
}

