package routes

import (
	"learn_api/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	eV1 := e.Group("api/v1/")
	eV1.GET("users", controllers.GetUserController)
	eV1.POST("users/login", controllers.LoginController)
	eV1.POST("users/register", controllers.RegisterController)
	eV1.GET("users/:userId", controllers.DetailUserController)
	e.Start(":8000")
	return e
}
