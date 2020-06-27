package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/joho/godotenv"

    "github.com/go-pg/pg/v10"
    //"github.com/go-pg/pg/v10/orm"
    "encoding/json"

    "log"
    "time"
    "os"
    "fmt"
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

    if order.ID != 0 {
        httpError := echo.ErrBadRequest 
        httpError.Message = "ID should not be set"
        return httpError
    }

    err := db.Insert(order)
    if err != nil {
        log.Printf("%v\n", err)
        httpError := echo.ErrBadRequest 
        httpError.Message = fmt.Sprintf("%s", err)
        return httpError
    }

    orderJson, err := json.Marshal(order)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, string(orderJson))
}

type Order struct {
    ID int `pg:"id,pk" json:"id"`
    CreatedAt time.Time `pg:"created_at" pg:"default:now()" json:"created_at"`
}
