package controller

import "github.com/labstack/echo/v4"

func getEntries(c echo.Context) error {
    return nil
}

func handleEntry(g *echo.Group) {
    r := g.Group("/entry")
    r.GET("", getEntries)
}

