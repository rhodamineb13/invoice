package handler

import (
	"invoice/common/dto"
	"invoice/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(ord service.OrderService) *OrderHandler {
	return &OrderHandler{
		ord,
	}
}

func (o *OrderHandler) DisplayOrders(c *gin.Context) {
	invIDString := c.Param("id")

	invID, err := strconv.Atoi(invIDString)
	if err != nil {
		_ = c.Error(err)
		return
	}

	lists, err := o.orderService.GetAllOrders(c, invID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (o *OrderHandler) AddNewOrder(c *gin.Context) {
	var ord *dto.OrdersDTO
	invIDString := c.Param("id")

	invID, err := strconv.Atoi(invIDString)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := c.ShouldBindJSON(&ord); err != nil {
		_ = c.Error(err)
		return
	}

	err = o.orderService.CreateNewOrders(c, invID, ord)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order added",
	})
}

func (o *OrderHandler) EditOrder(c *gin.Context) {
	var ord *dto.OrdersDTO
	invIDString := c.Param("id")

	invID, err := strconv.Atoi(invIDString)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := c.ShouldBindJSON(&ord); err != nil {
		_ = c.Error(err)
		return
	}

	err = o.orderService.CreateNewOrders(c, invID, ord)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order added",
	})
}

func (o *OrderHandler) DeleteOrder(c *gin.Context) {
	invIDString := c.Param("orderID")

	orderID, err := strconv.Atoi(invIDString)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = o.orderService.DeleteOrder(c, orderID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order deleted",
	})
}