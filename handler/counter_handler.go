package handler

import (
	"counter-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CounterHandler interface {
	GetNumber(ctx echo.Context)(error)
	IncrementNumber(ctx echo.Context)(error)
	DecrementNumber(ctx echo.Context)(error)
}

type counterHandler struct {
	service service.CounterService
}

func NewCounterHandler(e *echo.Echo, service service.CounterService) CounterHandler {
	handler := &counterHandler{service: service}
	e.GET("/", handler.GetNumber)
	e.PUT("/increment",handler.IncrementNumber)
	e.PUT("/decrement",handler.DecrementNumber)

	return handler
}

func (c *counterHandler) GetNumber(ctx echo.Context) error{
	status := http.StatusOK

	number, err := c.service.GetNumber()

	if err != nil {
		status = http.StatusInternalServerError
		return ctx.JSON(status, nil)
	}

	return ctx.JSON(status, number)
}

func (c *counterHandler) IncrementNumber(ctx echo.Context) error {
	status := http.StatusOK

	number, err := c.service.IncrementNumber()

	if err != nil {
		status = http.StatusInternalServerError
		return ctx.JSON(status, nil)
	}

	return ctx.JSON(status, number)
}

func (c *counterHandler) DecrementNumber(ctx echo.Context) error {
	status := http.StatusOK

	number, err := c.service.DecrementNumber()

	if err != nil {
		status = http.StatusInternalServerError
		return ctx.JSON(status, nil)
	}

	return ctx.JSON(status, number)
}
