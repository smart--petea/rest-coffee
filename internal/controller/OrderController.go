package controller

import (
    "strconv"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
    "github.com/smart--petea/rest-coffee/internal/entity"
)

type OrderController struct {
    *BaseController
}

func (orderController *OrderController) Get(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return orderController.HttpError(echo.ErrNotFound, err)
    }

    order := &entity.Order{ID: id}
    err = orderController.Db.
        Model(order).
        Relation("Items").
        Where("id = ?", id).
        Select()

    if err != nil {
        return orderController.HttpError(echo.ErrBadRequest, err)
    }

    return orderController.Response(c, order)
}

type postOrderBodyType struct {
    Items []*entity.OrderItem `json:"items"`
}

func (orderController *OrderController) Post(c echo.Context) error {
    tx, err := orderController.Db.Begin()
    if err != nil {
        return orderController.HttpError(echo.ErrInternalServerError, err)
    }

    createOrderBody := new(postOrderBodyType)
    if err := c.Bind(createOrderBody); err != nil {
        return orderController.HttpError(echo.ErrBadRequest, err)
    }

    order := new(entity.Order)
    err = tx.Insert(order)
    if err != nil {
        return orderController.HttpError(echo.ErrBadRequest, err)
    }

    order.Items = make([]*entity.OrderItem, 0, len(createOrderBody.Items))

    for _, orderItem := range createOrderBody.Items {
        orderItem.OrderId = order.ID
        err := tx.Insert(orderItem)
        if err != nil {
            tx.Rollback()

            return orderController.HttpError(echo.ErrBadRequest, err)
        }

        order.Items = append(order.Items, orderItem)
    }

    tx.Commit()

    return orderController.Response(c, order)
}

func NewOrderController(Db *pg.DB) *OrderController {
    return &OrderController{
        BaseController: &BaseController{Db: Db},
    }
}
