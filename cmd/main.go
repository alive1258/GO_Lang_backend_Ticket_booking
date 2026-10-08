package main

import (
	"fmt"
	"goticket/internal/config"
	"goticket/internal/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"gorm.io/gorm"
)


type User struct {
    gorm.Model

    Name string `json:"name" validate:"required" gorm:"type:varchar(100);not null"`

    Email string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique;not null"`

    Password string `json:"password" validate:"required,min=6" gorm:"type:varchar(100);not null"`
}

type CustomValidator struct {
    validator *validator.Validate
}



func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}



func main() {
   cfg :=  config.LoadEnv()

   db := config.ConnectDatabase(cfg)

    db.AutoMigrate(&User{})

    e := echo.New()
    e.Use(middleware.RequestLogger())
  

 
	

    e.GET("/", func(c *echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello, World!",
        })
    })

    e.Validator = &CustomValidator{validator: validator.New(),}
    
    // user route register
    user.RegisterRouters(e,db)


 port :=fmt.Sprintf(":%s", cfg.Port)
    if err := e.Start(port); err != nil {
        e.Logger.Error("failed to start server", "error", err)
    }
}