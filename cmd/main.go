package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/joho/godotenv"

    "log"

    "github.com/smart--petea/rest-coffee/internal/controller"
)

func main() {
    err := godotenv.Load("./configs/database.env")
    if err != nil {
        log.Fatal("Error loading .env. file")
    }

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    var orderController controller.Order
    e.POST("/order", orderController.Post)
    e.GET("/order/:id", orderController.Get)

    var goodController controller.Good
    e.POST("/good", goodController.Post)

    e.Logger.Fatal(e.Start(":9991")) //todo put the port in env. Don't forget about docker-compose.yml
}
