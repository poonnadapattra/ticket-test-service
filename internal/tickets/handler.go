package tickets

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetTicket(c echo.Context) error
}

type handler struct {
	service Service
}

func Newhandler(service Service) Handler {
	return handler{service}
}

func (h handler) GetTicket(c echo.Context) error {

	res, err := h.service.GetTicket()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
