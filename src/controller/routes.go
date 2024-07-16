package controller

import "github.com/labstack/echo/v4"

func HandlePaths(e *echo.Echo) {
    r := handleAuth(e)

    handleCollection(r)
    handleEntry(r)
    handleTag(r)
    handleCast(r)
    handleThumb(r)

    // media has its own auth
    handleMedia(e)
}

