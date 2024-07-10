package controller

import "github.com/labstack/echo/v4"

func HandlePaths(e *echo.Echo) {
    e.GET("/test", func(c echo.Context) error {
        return c.String(200, "test")
    })

    r := handleAuth(e)
    _ = r
}
