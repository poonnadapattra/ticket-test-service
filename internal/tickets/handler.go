package tickets

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/poonnadapattra/ticket-test-service/internal/entity"
)

type Handler interface {
	GetTicketCount(c echo.Context) error
	GetTicket(c echo.Context) error
	CreateTicket(c echo.Context) error
	UpdateTicket(c echo.Context) error
	DeleteTicket(c echo.Context) error
}

type handler struct {
	service Service
}

func Newhandler(service Service) Handler {
	return handler{service}
}

func (h handler) GetTicketCount(c echo.Context) error {

	res, err := h.service.GetTicketCount()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Code:    http.StatusOK,
		Massage: http.StatusText(http.StatusOK),
		Data:    res,
	})
}

func (h handler) GetTicket(c echo.Context) error {

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 0
		log.Error(err, "parsing: page params")
	}
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = 0
		log.Error(err, "parsing: size params")
	}
	req := ReqTicket{}
	req.Status = c.QueryParam("status")
	req.Pagging.Page = page
	req.Pagging.Size = size

	res, err := h.service.GetTicket(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Code:    http.StatusOK,
		Massage: http.StatusText(http.StatusOK),
		Data:    res,
	})
}

func (h handler) CreateTicket(c echo.Context) error {

	var req Tickets
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.CreateTicket(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Code:    http.StatusOK,
		Massage: http.StatusText(http.StatusOK),
	})
}

func (h handler) UpdateTicket(c echo.Context) error {
	var req Tickets
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.UpdateTicket(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Code:    http.StatusOK,
		Massage: http.StatusText(http.StatusOK)})
}

func (h handler) DeleteTicket(c echo.Context) error {

	var req Tickets
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.DeleteTicket(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Code:    http.StatusOK,
		Massage: http.StatusText(http.StatusOK)})
}
