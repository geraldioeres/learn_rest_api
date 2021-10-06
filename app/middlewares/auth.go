package middlewares

import (
	"learn_api/controllers"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTMyClaims struct {
	UserId int
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT      string
	ExipesDuration int
}

func (config *ConfigJWT) Init() middleware.JWTConfig{
	return middleware.JWTConfig{
		Claims: &JWTMyClaims{},
		SigningKey: []byte(config.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(
			func(e error, c echo.Context) error {
				return controllers.NewErrorResponse(c, http.StatusForbidden, e)
			},
		),
	}
}