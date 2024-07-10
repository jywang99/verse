package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"jy.org/verse/src/config"
	"jy.org/verse/src/service"
)

var authCfg = config.Config.Auth

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
    Id int `json:"id"`
	jwt.RegisteredClaims
}

var signKey = []byte(authCfg.Secret)

type Login struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func login(c echo.Context) error {
    l := new(Login)
    if err := c.Bind(l); err != nil {
        return err
    }
    // TODO validation
    if l.Email == "" || l.Password == "" {
        return echo.ErrUnauthorized
    }

    user, ok := service.Authenticate(l.Email, l.Password)
    if !ok {
        return echo.ErrUnauthorized
    }

	// Set custom claims
	claims := &jwtCustomClaims{
        user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(signKey)
    if err != nil {
        return handleError(c, err)
    }

	return c.JSONPretty(http.StatusOK, echo.Map{
        "username": user.Username,
		"token": t,
	}, "  ")
}

type Register struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Name string `json:"name"`
}

func register(c echo.Context) error {
    // TODO validator
    r := new(Register)
    if err := c.Bind(r); err != nil {
        return err
    }

    // TODO validation
    reg := service.Register{
        Username: r.Username,
        Email: r.Email,
        Password: r.Password,
        Name: r.Name,
    }
    err := reg.Register()
    if err != nil {
        return handleError(c, err)
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "success",
    })
}

// Parse claims from JWT token to retrieve user info
func GetClaims(c echo.Context) *jwtCustomClaims {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*jwtCustomClaims)
    return claims
}

func handleAuth(e *echo.Echo) *echo.Group {
    e.POST("/login", login)
    e.POST("/regist", register)

    r := e.Group("/my")
    config := echojwt.Config{
        NewClaimsFunc: func(c echo.Context) jwt.Claims {
            return new(jwtCustomClaims)
        },
        SigningKey: signKey,
    }
    r.Use(echojwt.WithConfig(config))

    return r
}

