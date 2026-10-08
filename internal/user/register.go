package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRouters(e *echo.Echo, db *gorm.DB) {
	userRepository := NewRepository(db)

	userService := NewService(userRepository)

	userHandler := NewHandler(userService)

	e.POST("/users", userHandler.CreateUser)
}