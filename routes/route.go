package routes

import (
	"learn_api/controllers"
	"learn_api/controllers/twitters"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	eV1 := e.Group("api/v1/")
	eV1.GET("users", controllers.GetUserController)
	eV1.POST("users/login", controllers.LoginController)
	eV1.POST("users/register", controllers.RegisterController)
	eV1.GET("users/:userId", controllers.DetailUserController)

	eV1.GET("tweet", twitters.GetTwitterController)

	eV1.POST("wish", controllers.WishRegister)
	eV1.POST("products", controllers.ProductRegister)
	eV1.GET("get", controllers.GetWish)
	eV1.GET("products", controllers.GetProduct)
	return e
}
