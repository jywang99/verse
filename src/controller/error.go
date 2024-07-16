package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"jy.org/verse/src/except"
)

func handleError(c echo.Context, err error) error {
    body := map[string]string{"error": err.Error()}
    handledErr, ok := err.(except.HandledError)
    if !ok {
        log.Printf("Unhandled error: %v", handledErr)
        return c.JSON(http.StatusInternalServerError, body)
    }
    switch handledErr.Type {
        case except.AuthErr:
            return c.JSON(http.StatusUnauthorized, body)
        case except.NotFoundErr:
            return c.JSON(http.StatusNotFound, body)
        case except.BadRequestErr:
            return c.JSON(http.StatusBadRequest, body)
        case except.ForbiddenErr:
            return c.JSON(http.StatusForbidden, body)
        case except.DbErr:
            return c.JSON(http.StatusInternalServerError, body)
        default:
            log.Printf("Unhandled error: %v", handledErr)
            return c.JSON(http.StatusInternalServerError, body)
    }
}

