package controller

import (
    "github.com/labstack/echo/v4"
    "fmt"
    "encoding/json"
    "net/http"
)

type BaseController struct {
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
