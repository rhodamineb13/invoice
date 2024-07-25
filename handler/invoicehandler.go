package handler

import (
	"invoice/common/dto"
	"invoice/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService service.InvoiceService
}

func NewInvoiceHandler(inv service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		inv,
	}
}

func (i *InvoiceHandler) GetAllInvoice(c *gin.Context) {
	lists, err := i.invoiceService.InvoiceIndex(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (i *InvoiceHandler) GetInvoiceByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		_ = c.Error(err)
		return
	}

	inv, err := i.invoiceService.InvoiceByID(c, id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, inv)
}

func (i *InvoiceHandler) NewInvoice(c *gin.Context) {
	var inv *dto.InvoiceDetailDTO

	if err := c.ShouldBindJSON(&inv); err != nil {
		_ = c.Error(err)
		return
	}

	if err := i.invoiceService.AddInvoice(c, inv); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "invoice successfully created",
	})

}

func (i *InvoiceHandler) EditInvoice(c *gin.Context) {
	var inv *dto.InvoiceDetailDTO

	if err := c.ShouldBindJSON(&inv); err != nil {
		_ = c.Error(err)
		return
	}

	if err := i.invoiceService.EditInvoice(c, inv); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "invoice successfully edited",
	})

}
