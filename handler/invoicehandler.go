package handler

import (
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService InvoiceService
}

func newInvoiceHandler(inv InvoiceService) *invoiceHandler {
	return &InvoiceHandler{
		inv,
	}
}

func (i *InvoiceHandler) AllInvoice(c *gin.Context) {

}

func (i *InvoiceHandler) GetInvoiceByID(c *gin.Context) {

}
