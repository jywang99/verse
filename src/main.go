package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jy.org/verse/src/config"
	cs "jy.org/verse/src/constant"
	"jy.org/verse/src/controller"
	"jy.org/verse/src/logging"
)

var cfg = config.Config
var logger = logging.Logger

func main() {
    e := echo.New()
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: cfg.Server.AllowedOrigins,
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowCredentials},
        AllowCredentials: true,
        ExposeHeaders: []string{cs.ContentRange},
    }))

    e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
        LogStatus:   true,
        LogURI:      true,
        LogError:    true,
        LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
            if v.Error == nil {
                logger.INFO.Printf("uri=%s status=%d\n", v.URI, v.Status)
            } else {
                logger.ERROR.Printf("uri=%s status=%d error=%v\n", v.URI, v.Status, v.Error)
            }
            return nil
        },
    }))
    e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
        LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
            logger.ERROR.Printf("error=%v stack=%s\n", err, stack)
            return err
        },
    }))

    controller.HandlePaths(e)

    e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", cfg.Server.Port)))
}

