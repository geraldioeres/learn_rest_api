package controllers

import (
	"learn_api/configs"
	"learn_api/models/product"
	"learn_api/models/response"
	"learn_api/models/wish"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProduct(c echo.Context) error {
	product := []product.Product{}

	result := configs.DB.Preload("Wish").Find(&product)

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
		Data:    result,
	})
}

func GetWish(c echo.Context) error {
	wish := []wish.Wish{}

	result := configs.DB.Preload("Product").Find(&wish)

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
		Data:    wish,
	})
}

func WishRegister(c echo.Context) error {
	var wishReg wish.WishRegister
	c.Bind(&wishReg)

	var wishDB wish.Wish
	wishDB.Name = wishReg.Name

	result := configs.DB.Create(&wishDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data wish ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Register Wish",
		Data:    wishDB,
	})
}

func ProductRegister(c echo.Context) error {
	var prodReg product.ProductRegister
	c.Bind(&prodReg)

	var prodDB product.Product
	prodDB.Name = prodReg.Name
	prodDB.WishId = prodReg.WishId

	result := configs.DB.Create(&prodDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data product ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Register Product",
		Data:    prodDB,
	})
}
