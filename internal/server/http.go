package server

import (
	"fmt"
	"goticket/internal/config"
	"os/user"

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
	db.AutoMigrate(&user.User{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())

	// e.GET("/health", func(c *echo.Context) error {
	// 	return c.String(http.StatusOK, "ok")
	// })

	// user routes
	// user.RegisterRoutes(e, db)

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}