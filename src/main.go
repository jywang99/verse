package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jy.org/verse/src/config"
	"jy.org/verse/src/controller"
	"jy.org/verse/src/db"
	"jy.org/verse/src/logging"
)

var cfg = config.Config
var logger = logging.Logger

func main() {
    config.Init("conf/config.yml")
    logging.Init()
    db.Init()

    e := echo.New()
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.Logger()) // TODO to file
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: cfg.Server.AllowedOrigins,
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
    }))

    controller.HandlePaths(e)

    e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", cfg.Server.Port)))
}

