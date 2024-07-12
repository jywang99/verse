package controller

import "github.com/labstack/echo/v4"

func HandlePaths(e *echo.Echo) {
    r := handleAuth(e)
    handleCollections(r)
}

