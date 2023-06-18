package contacts

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/poonnadapattra/ticket-test-service/internal/entity"
)

type Handler interface {
	GetContact(c echo.Context) error
}

type handler struct {
	service Service
}

func Newhandler(service Service) Handler {
	return handler{service}
}

func (h handler) GetContact(c echo.Context) error {

	res, err := h.service.GetContact()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Code:    http.StatusOK,
		Massage: http.StatusText(http.StatusOK),
		Data:    res,
	})
}
