package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type BaseResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func main() {
	e := echo.New()

	e.GET("v1/users", GetUserController)
	e.Start(":8000")
}

func GetUserController(c echo.Context) error {
	user := User{"Alterra", "alterrra@gmail.com", "malang"}
	return c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Message: "Berhasil",
		Data: user,
	})
}
