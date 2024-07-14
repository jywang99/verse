package controller

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"jy.org/verse/src/except"
)

func parseIdParam(c echo.Context, idStr string) (int, error) {
    id, err := strconv.Atoi(c.Param(idStr))
    if err != nil {
        return 0, except.NewHandledError(except.BadRequestErr, fmt.Sprintf("Invalid id: %s", idStr))
    }
    if id < 1 {
        return 0, except.NewHandledError(except.BadRequestErr, fmt.Sprintf("Invalid id: %s", idStr))
    }
    return id, nil
}

