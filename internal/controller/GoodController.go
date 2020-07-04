package controller

import (
    "github.com/labstack/echo/v4"
    "github.com/smart--petea/rest-coffee/internal/entity"
    "github.com/go-pg/pg/v10"
    "strconv"
    "errors"
)

type GoodController struct {
    *BaseController
}

func (goodController *GoodController) Post(c echo.Context) error {
    good := new(entity.Good)
    if err := c.Bind(good); err != nil {
        return goodController.HttpError(echo.ErrNotFound, err)
    }

    if good.ID != 0 {
        err := errors.New("Id should not be setted")
        return goodController.HttpError(echo.ErrBadRequest, err)
    }

    err := goodController.Db.Insert(good)
    if err != nil {
        return goodController.HttpError(echo.ErrBadRequest, err)
    }

    return goodController.Response(c, good)
}

func (goodController *GoodController) Get(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return goodController.HttpError(echo.ErrBadRequest, err)
    }

    good := &entity.Good{ID: id}
    err = goodController.Db.
        Model(good).
        Where("id = ?", id).
        Select()
    if err != nil {
        return goodController.HttpError(echo.ErrInternalServerError, err)
    }

    return goodController.Response(c, good)
}

func (goodController *GoodController) Routing() {
    goodController.POST("/good", goodController.Post)
    goodController.GET("/good/:id", goodController.Get)
}

func NewGoodController(Db *pg.DB, e *echo.Echo) *GoodController {
    goodController := GoodController{
        BaseController: NewBaseController(Db, e),
    }

    goodController.Routing()

    return &goodController
}
