package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/joho/godotenv"
    "github.com/smart--petea/rest-coffee/internal/controller"
    "github.com/smart--petea/rest-coffee/internal/helper"
)

func main() {
    e := echo.New()

    err := godotenv.Load("./configs/database.env")
    if err != nil {
        e.Logger.Fatal("Error loading .env. file")
    }


    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    dbConnection := helper.GetDb()
    controller.NewOrderController(dbConnection, e)
    controller.NewGoodController(dbConnection, e)

    e.Logger.Fatal(e.Start(":9991")) //todo put the port in env. Don't forget about docker-compose.yml
}
