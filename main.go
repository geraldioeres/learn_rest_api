package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Address   string         `json:"address"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type BaseResponse struct {
	Code    int
	Message string
	Data    interface{}
}

var DB *gorm.DB

func InitDB() {
	dsn := "root:diosql@tcp(127.0.0.1:3306)/learn_api"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed to connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&User{})
}

func main() {
	InitDB()
	e := echo.New()
	eV1 := e.Group("v1/")
	eV1.GET("users", GetUserController)
	eV1.POST("users/login", LoginController)
	eV1.POST("users/register", RegisterController)
	eV1.GET("users/:userId", DetailUserController)
	e.Start(":8000")
}

func RegisterController(c echo.Context) error {
	var userRegister UserRegister
	c.Bind(&userRegister)

	//Validation
	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama masih kosong",
			Data:    nil,
		})
	}
	//Etc

	var userDB User
	userDB.Name = userRegister.Name
	userDB.Password = userRegister.Password
	userDB.Address = userRegister.Address
	userDB.Email = userRegister.Email

	result := DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data user ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Register",
		Data:    userDB,
	})
}

func LoginController(c echo.Context) error {
	userLogin := UserLogin{}
	c.Bind(&userLogin)

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    userLogin,
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
	// name := c.QueryParam("name")
	// address := c.QueryParam("address")
	// bisnis
	user := User{}

	// if name == "" {
	// 	user = User{1, "Alterra", "alterrra@gmail.com", "malang", ""}
	// } else {
	// 	user = User{1, name, "alterrra@gmail.com", address, ""}
	// }

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    user,
	})
}
