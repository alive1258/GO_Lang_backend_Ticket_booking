package main

import (
	"goticket/internal/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
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
dsn := "postgresql://neondb_owner:npg_vJtVsKu08cnx@ep-young-bar-b5s5zx7s-pooler.c-7.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    TranslateError: true,
})



if err != nil {
    panic("failed to connect database")
}else{
    println("Database connection established successfully.")
}

db.AutoMigrate(&User{})

    e := echo.New()

    e.Use(middleware.RequestLogger())
    e.Use(middleware.Recover())

 
	

    e.GET("/", func(c *echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Hello, World!",
        })
    })

    e.Validator = &CustomValidator{validator: validator.New(),}
    
    // user route register
    user.RegisterRouters(e,db)



    if err := e.Start(":8080"); err != nil {
        e.Logger.Error("failed to start server", "error", err)
    }
}