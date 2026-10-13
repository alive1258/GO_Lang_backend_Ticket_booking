package user

import (
	"goticket/internal/auth"
	"goticket/internal/config"
	"goticket/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRouters(e *echo.Echo, db *gorm.DB,cfg *config.Config ) {
	userRepository := NewRepository(db)
    jwtService := auth.NewJWTService(cfg.JwtSecret)
	userService := NewService(userRepository,jwtService)

	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/me", userHandler.GetMe, middlewares.AuthMiddleware(jwtService)) // protected route
}