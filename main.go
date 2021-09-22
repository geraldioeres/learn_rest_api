package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id      int    `json:"id"`
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
	eV1 := e.Group("v1/")
	eV1.GET("users", GetUserController)
	eV1.POST("users/login", LoginController)
	eV1.GET("users/:userId", DetailUserController)
	e.Start(":8000")
}

func LoginController(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    User{Id: 1, Email: email, Address: password},
	})

}

func DetailUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    User{Id: userId},
	})
}

func GetUserController(c echo.Context) error {
	name := c.QueryParam("name")
	address := c.QueryParam("address")
	// bisnis
	user := User{}

	if name == "" {
		user = User{1, "Alterra", "alterrra@gmail.com", "malang"}
	} else {
		user = User{1, name, "alterrra@gmail.com", address}
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    user,
	})
}
