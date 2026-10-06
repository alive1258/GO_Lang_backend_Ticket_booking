package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type User struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type CustomValidator struct {
    validator *validator.Validate
}



func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func executeSomeBusinessLogic(user User) {
    // save user to database etc.
}

func main() {

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


    e := echo.New()

    e.Use(middleware.RequestLogger())
    e.Use(middleware.Recover())

    e.Validator = &CustomValidator{
        validator: validator.New(),
    }
	

    e.GET("/", func(c *echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello, World!",
        })
    })

    e.POST("/users", func(c *echo.Context) error {
        u := new(User)

        if err := c.Bind(u); err != nil {
            return c.String(http.StatusBadRequest, "bad request")
        }

        if err := c.Validate(u); err != nil {
            return c.String(http.StatusBadRequest, err.Error())
        }

        executeSomeBusinessLogic(*u)

        return c.JSON(http.StatusOK, u)
    })

    if err := e.Start(":8080"); err != nil {
        e.Logger.Error("failed to start server", "error", err)
    }
}