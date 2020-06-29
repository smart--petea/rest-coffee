package controller

import (
    "github.com/labstack/echo/v4"

    "github.com/smart--petea/rest-coffee/internal/entity"
    "github.com/smart--petea/rest-coffee/internal/helper"

    "log"
    "net/http"
    "encoding/json"
    "fmt"
    "strconv"
)

type Good struct {
    Base
}

func (*Good) Post(c echo.Context) error {
    db := helper.GetDb()
    defer db.Close()

    good := new(entity.Good)
    if err := c.Bind(good); err != nil {
        return nil
    }

    if good.ID != 0 {
        httpError := echo.ErrBadRequest 
        httpError.Message = "ID should not be set"
        return httpError
    }

    err := db.Insert(good)
    if err != nil {
        log.Printf("%v\n", err)
        httpError := echo.ErrBadRequest 
        httpError.Message = fmt.Sprintf("%s", err)
        return httpError
    }

    goodJson, err := json.Marshal(good)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, string(goodJson))
}

func (*Good) Get(c echo.Context) error {
    db := helper.GetDb()
    defer db.Close()

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("%v\n", err)
        httpError := echo.ErrBadRequest 
        httpError.Message = fmt.Sprintf("%s", err)
        return httpError
    }

    good := &entity.Good{ID: id}
    err = db.
        Model(good).
        Where("id = ?", id).
        Select()

    if err != nil {
        log.Printf("%v\n", err)
        return echo.ErrNotFound
    }

    goodJson, err := json.Marshal(good)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, string(goodJson))
}
