package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/joho/godotenv"

    "github.com/go-pg/pg/v10"
    //"github.com/go-pg/pg/v10/orm"

    "log"
    "time"
    "os"
)

func main() {

    err := godotenv.Load("./configs/database.env")
    if err != nil {
        log.Fatal("Error loading .env. file")
    }

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.POST("/order", createOrder)

    e.Logger.Fatal(e.Start(":9991"))
}

type OrderItem struct {
}

func createOrder(c echo.Context) error {
    db := pg.Connect(&pg.Options{
        Addr: "database:5432",
        User:  os.Getenv("POSTGRES_USER"),
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Database: os.Getenv("POSTGRES_DB"),
    })
    defer db.Close()

    order := new(Order)
    if err := c.Bind(order); err != nil {
        return nil
    }

    err := db.Insert(order)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, "Order created")
}

type Order struct {
    Id int `pg:"id" json:"id"`
    CreatedAt time.Time `pg:"created_at" pg:"default:now()" json:"created_at"`
}
