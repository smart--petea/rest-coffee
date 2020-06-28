package controller

import (
    "strconv"

    "github.com/labstack/echo/v4"
    "fmt"
    "log"
    "encoding/json"
    "net/http"

    "github.com/smart--petea/rest-coffee/internal/entity"
    "github.com/smart--petea/rest-coffee/internal/helper"
)

type Order struct {}

func (*Order) Get(c echo.Context) error {
    db := helper.GetDb()
    defer db.Close()

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("%v\n", err)
        httpError := echo.ErrBadRequest 
        httpError.Message = fmt.Sprintf("%s", err)
        return httpError
    }

    order := &entity.Order{ID: id}
    err = db.
        Model(order).
        Relation("Items").
        Where("id = ?", id).
        Select()

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

type postOrderBodyType struct {
    Items []*entity.OrderItem `json:"items"`
}

func (*Order) Post(c echo.Context) error {
    db := helper.GetDb()
    defer db.Close()

    log.Printf("\n45\n")

    createOrderBody := new(postOrderBodyType)
    if err := c.Bind(createOrderBody); err != nil {
        log.Printf("%v\n", err)
        httpError := echo.ErrBadRequest 
        httpError.Message = fmt.Sprintf("%s", err)
        return httpError
    }

    order := new(entity.Order)
    err := db.Insert(order)
    if err != nil {
        log.Printf("%v\n", err)
        httpError := echo.ErrBadRequest 
        httpError.Message = fmt.Sprintf("%s", err)
        return httpError
    }

    order.Items = make([]*entity.OrderItem, 0, len(createOrderBody.Items))

    for _, orderItem := range createOrderBody.Items {
        orderItem.OrderId = order.ID
        err := db.Insert(orderItem)
        if err != nil {
            log.Printf("%v\n", err)
            httpError := echo.ErrBadRequest 
            httpError.Message = fmt.Sprintf("%s", err)
            return httpError
        }

        order.Items = append(order.Items, orderItem)
    }

    //todo transaction
    //todo array of order items

    orderJson, err := json.Marshal(order)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, string(orderJson))
}