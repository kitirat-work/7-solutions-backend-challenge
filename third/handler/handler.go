package handler

import (
	"piefiredire/service"

	"github.com/labstack/echo/v4"
)

type PieFireDireHandler interface {
	BeefSummary(c echo.Context) error
}

type Handler struct {
	srv service.PieFireDireService
}

func NewHandler(
	srv service.PieFireDireService,
) PieFireDireHandler {
	return &Handler{
		srv: srv,
	}
}

// BeefSummary implements PieFireDireHandler.
func (h *Handler) BeefSummary(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.srv.BeefSummary(ctx)
	if err != nil {
		return err
	}
	return c.JSON(200, result)
}
