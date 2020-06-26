package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/go-pg/pg/v10"
    //"github.com/go-pg/pg/v10/orm"
)

func main() {
    db := pg.Connect(&pg.Options{
        Addr: ":9992",
        User:  "postgres",
        Password: "postgres",
    })
    defer db.Close()




    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/order", createOrder)

    e.Logger.Fatal(e.Start(":9991"))
}

type Order struct {
}

type OrderItem struct {
}

func createOrder(c echo.Context) error {
    return c.String(http.StatusOK, "Order created")
}
