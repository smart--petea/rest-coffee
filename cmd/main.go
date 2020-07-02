package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/joho/godotenv"

    "log"

    "github.com/smart--petea/rest-coffee/internal/controller"
    "github.com/smart--petea/rest-coffee/internal/helper"
)

func main() {
    err := godotenv.Load("./configs/database.env")
    if err != nil {
        log.Fatal("Error loading .env. file")
    }

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    dbConnection := helper.GetDb()

    orderController := controller.Order{controller.BaseController{Db: dbConnection}}
    e.POST("/order", orderController.Post)
    e.GET("/order/:id", orderController.Get)

    goodController := controller.Good{controller.BaseController{Db: dbConnection}}
    e.POST("/good", goodController.Post)
    e.GET("/good/:id", goodController.Get)

    e.Logger.Fatal(e.Start(":9991")) //todo put the port in env. Don't forget about docker-compose.yml
}
