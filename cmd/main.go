package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/joho/godotenv"
    "github.com/smart--petea/rest-coffee/internal/controller"
    "github.com/smart--petea/rest-coffee/internal/helper"

    "log"
    "os"
)

func main() {

    err := godotenv.Load("./configs/database.env")
    if err != nil {
        log.Fatal("Error loading .env. file")
    }

    e := echo.New()

    f, err := os.OpenFile("log/api.log", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Cannot open log file, (%s)", err.Error())
    }
    defer f.Close()

    loggerConfig := middleware.DefaultLoggerConfig
    loggerConfig.Output = f
    e.Use(middleware.LoggerWithConfig(loggerConfig)) //to file

    e.Use(middleware.Logger()) //to stdout
    e.Use(middleware.Recover())

    dbConnection := helper.GetDb()
    controller.NewOrderController(dbConnection, e)
    controller.NewGoodController(dbConnection, e)

    e.Logger.Fatal(e.Start(":9991")) //todo put the port in env. Don't forget about docker-compose.yml
}
