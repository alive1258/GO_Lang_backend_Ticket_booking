package server

import (
	"fmt"

	"goticket/internal/config"
	"goticket/internal/event"
	"goticket/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(db *gorm.DB, cfg *config.Config) {
	// Auto migrate
	if err := db.AutoMigrate(&user.User{}); err != nil {
		panic("failed to migrate database")
	}

	// Create Echo
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// Middleware
	e.Use(middleware.RequestLogger())

	// Register routes
	user.RegisterRouters(e, db)
	event.RegisterRoutes(e, db)

	// Port
	port := fmt.Sprintf(":%s", cfg.Port)

	// Start server
	if err := e.Start(port); err != nil {
		e.Logger.Error(
			"failed to start server",
			"error",
			err,
		)
	}
}