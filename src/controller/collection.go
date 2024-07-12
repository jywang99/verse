package controller

import (
	"github.com/labstack/echo/v4"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/service"
)

func getCollections(c echo.Context) error {
    cg := e.DefaultCollectionGet()
    if err := c.Bind(cg); err != nil {
        return err
    }

    if err := cg.Validate(); err != nil {
        return handleError(c, err)
    }

    cols, err := service.GetCollections(cg)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, cols)
}

func handleCollections(g *echo.Group) {
    r := g.Group("/collections")
    r.GET("", getCollections)
}

