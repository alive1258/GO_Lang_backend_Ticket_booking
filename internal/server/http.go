package server

import (
	"fmt"
	"goticket/internal/config"
	"goticket/internal/domain/booking"
	"goticket/internal/domain/event"
	"goticket/internal/domain/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Start(db *gorm.DB, cfg *config.Config) {
	// Auto migrate
   db.AutoMigrate(&user.User{}, &event.Event{}, &booking.Booking{})

	// Create Echo
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// Middleware
	e.Use(middleware.RequestLogger())

		e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "running move")
	})

	// Register routes
	user.RegisterRouters(e, db,cfg)
	event.RegisterRoutes(e, db)
	booking.RegisterRoutes(e, db,cfg)

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