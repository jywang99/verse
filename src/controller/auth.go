package controller

import (
    "net/http"
    "time"

    "github.com/golang-jwt/jwt/v4"
    echojwt "github.com/labstack/echo-jwt"
    "github.com/labstack/echo/v4"
    e "jy.org/verse/src/entity"
    "jy.org/verse/src/service"
)

// jwtUserClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtUserClaims struct {
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

    user, err := service.Authenticate(l.Email, l.Password)
    if err != nil {
        return handleError(c, err)
    }

    // Set custom claims
    claims := &jwtUserClaims{
        user.Id,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * authCfg.TokenDuration)),
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
        "username": user.Name,
        "token": t,
    }, "  ")
}

type Register struct {
    Name string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func register(c echo.Context) error {
    // TODO validator
    r := new(Register)
    if err := c.Bind(r); err != nil {
	return err
    }

    user := e.RegistUser{
	Name: r.Name,
	Email: r.Email,
	Password: r.Password,
    }
    err := service.Register(user)
    if err != nil {
	return handleError(c, err)
    }

    return c.JSON(http.StatusOK, echo.Map{
	"message": "success",
    })
}

// Parse claims from JWT token to retrieve user info
func getClaims(c echo.Context) *jwtUserClaims {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*jwtUserClaims)
    return claims
}

func handleAuth(e *echo.Echo) *echo.Group {
    e.POST("/login", login)
    e.POST("/regist", register)

    r := e.Group("")
    config := echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	    return new(jwtUserClaims)
	},
	SigningKey: signKey,
    }
    r.Use(echojwt.WithConfig(config))

    return r
}

