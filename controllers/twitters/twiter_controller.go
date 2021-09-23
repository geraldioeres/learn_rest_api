package twitters

import (
	"learn_api/configs"
	"learn_api/models/response"
	"learn_api/models/twitters"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTwitterController(c echo.Context) error {
	twitters := []twitters.Tweet{}

	result := configs.DB.Preload("User").Find(&twitters)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika mendapatkan data twitter dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data twitter",
		Data:    twitters,
	})
}
