package middlewares

import (
	"learn_api/configs"
	"learn_api/models/users"

	"github.com/labstack/echo/v4"
)

func BasicAuth(email, password string, c echo.Context) (bool, error) {
	user := users.User{}
	result := configs.DB.First(&user, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
