package routes

import (
	"learn_api/constants"
	"learn_api/controllers"
	"learn_api/controllers/twitters"
	"learn_api/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.BodyDump(middlewares.Log))

	jwt := middleware.JWT([]byte(constants.SECRET_JWT))
	eV1 := e.Group("api/v1/")
	eV1.GET("users", controllers.GetUserController, jwt)
	eV1.POST("users/login", controllers.LoginController)
	eV1.POST("users/register", controllers.RegisterController)
	eV1.GET("users/:userId", controllers.DetailUserController)

	eV1.GET("tweet", twitters.GetTwitterController)
	return e
}
