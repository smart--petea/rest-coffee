package controller

import (
    "github.com/labstack/echo/v4"
    "fmt"
)

type Base struct {
}

func (base *Base) HttpError(httpError *echo.HTTPError, err error) *echo.HTTPError {
    httpError.Message = fmt.Sprintf("%s", err)
    return httpError
}
