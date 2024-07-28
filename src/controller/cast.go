package controller

import (
	"github.com/labstack/echo/v4"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
	"jy.org/verse/src/service"
)

func getCasts(c echo.Context) error {
    get := e.NewGetCasts()
    if err := c.Bind(&get); err != nil {
        return handleError(c, err)
    }

    if err := get.Validate(); err != nil {
        return handleError(c, err)
    }

    got, err := service.GetCasts(*get)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, *got)
}

func getCastsByIds(c echo.Context) error {
    var get e.GetByIds
    if err := c.Bind(&get); err != nil {
        return handleError(c, err)
    }
    if len(get.Ids) == 0 {
        return handleError(c, except.NewHandledError(except.BadRequestErr, "No cast ids provided"))
    }

    got, err := service.GetCastByIds(get.Ids)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, got)
}

func getCast(c echo.Context) error {
    id, err := parseIdParam(c, "id")
    if err != nil {
        return handleError(c, err)
    }

    got, err := service.GetCastById(id)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, *got)
}

func handleCast(g *echo.Group) {
    r := g.Group("/cast")
    r.POST("", getCasts)
    r.POST("/ids", getCastsByIds)
    r.GET("/:id", getCast)
}
