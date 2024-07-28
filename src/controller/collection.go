package controller

import (
	"github.com/labstack/echo/v4"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
	"jy.org/verse/src/service"
)

func getCollections(c echo.Context) error {
    cg := e.DefaultGetCollections()
    if err := c.Bind(cg); err != nil {
        return err
    }

    if err := cg.Validate(); err != nil {
        return handleError(c, err)
    }

    cols, err := service.GetCollection(*cg)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, cols)
}

func getCollectionsByIds(c echo.Context) error {
    var get e.GetByIds
    if err := c.Bind(&get); err != nil {
        return handleError(c, err)
    }
    if len(get.Ids) == 0 {
        return handleError(c, except.NewHandledError(except.BadRequestErr, "No collection ids provided"))
    }

    got, err := service.GetCollectionsByIds(get.Ids)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, got)
}

func handleCollection(g *echo.Group) {
    r := g.Group("/collection")
    r.POST("", getCollections)
    r.POST("/ids", getCollectionsByIds)
}

