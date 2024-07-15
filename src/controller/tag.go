package controller

import (
	"github.com/labstack/echo/v4"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/service"
)

func getTags(c echo.Context) error {
    get := e.NewGetTags()
    if err := c.Bind(&get); err != nil {
        return handleError(c, err)
    }

    if err := get.Validate(); err != nil {
        return handleError(c, err)
    }

    got, err := service.GetTags(*get)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, *got)
}

func getTag(c echo.Context) error {
    id, err := parseIdParam(c, "id")
    if err != nil {
        return handleError(c, err)
    }

    got, err := service.GetTagById(id)
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(200, *got)
}

func handleTag(g *echo.Group) {
    r := g.Group("/tag")
    r.POST("", getTags)
    r.GET("/:id", getTag)
}

