package controller

import (
    "github.com/labstack/echo/v4"
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/go-pg/pg/v10"
)

type BaseController struct {
    Db *pg.DB
    Echo *echo.Echo
}

func (baseController *BaseController) HttpError(httpError *echo.HTTPError, err error) *echo.HTTPError {
    httpError.Message = fmt.Sprintf("%s", err)
    return httpError
}

func (baseController *BaseController) Response(c echo.Context, obj interface{}) error {
    objJson, err := json.Marshal(obj)
    if err != nil {
        return baseController.HttpError(echo.ErrInternalServerError, err)
    }

    return c.String(http.StatusOK, string(objJson))
}

func (baseController *BaseController) POST(path string, h echo.HandlerFunc) {
    baseController.Echo.POST(path, h)
}

func (baseController *BaseController) GET(path string, h echo.HandlerFunc) {
    baseController.Echo.GET(path, h)
}

func NewBaseController(Db *pg.DB, Echo *echo.Echo) *BaseController {
    return &BaseController{
        Db: Db,
        Echo: Echo,
    }
}
