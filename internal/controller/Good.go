package controller

import (
    "github.com/labstack/echo/v4"

    "github.com/smart--petea/rest-coffee/internal/entity"
    "github.com/smart--petea/rest-coffee/internal/helper"

    "log"
    "net/http"
    "encoding/json"
    "strconv"
    "errors"
)

type Good struct {
    BaseController
}

func (goodController *Good) Post(c echo.Context) error {
    db := helper.GetDb()
    defer db.Close()

    good := new(entity.Good)
    if err := c.Bind(good); err != nil {
        return goodController.HttpError(echo.ErrNotFound, err)
    }

    if good.ID != 0 {
        err := errors.New("Id should not be setted")
        return goodController.HttpError(echo.ErrBadRequest, err)
    }

    err := db.Insert(good)
    if err != nil {
        return goodController.HttpError(echo.ErrBadRequest, err)
    }

    goodJson, err := json.Marshal(good)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, string(goodJson))
}

func (goodController *Good) Get(c echo.Context) error {
    db := helper.GetDb()
    defer db.Close()

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return goodController.HttpError(echo.ErrBadRequest, err)
    }

    good := &entity.Good{ID: id}
    err = db.
        Model(good).
        Where("id = ?", id).
        Select()
    if err != nil {
        return goodController.HttpError(echo.ErrInternalServerError, err)
    }

    goodJson, err := json.Marshal(good)
    if err != nil {
        log.Printf("%v\n", err)
        return nil
    }

    return c.String(http.StatusOK, string(goodJson))
}
