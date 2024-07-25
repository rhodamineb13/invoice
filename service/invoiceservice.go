package service

import (
	"context"
	"invoice/common/dto"
	"invoice/repository"
)

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

type InvoiceService interface {
	InvoiceIndex(context.Context) ([]dto.InvoiceLists, error)
	InvoiceByID(context.Context, int) (*dto.InvoiceDetail, error)
	AddInvoice(context.Context, *dto.InvoiceDetail) error
	EditInvoice(context.Context, *dto.InvoiceDetail) error
}

func NewInvoiceService(inv repository.InvoiceRepository) InvoiceService {
	return &invoiceService{
		inv,
	}
}

func (in *invoiceService) InvoiceIndex(ctx context.Context) ([]dto.InvoiceLists, error) {
	return nil, nil
}

func (in *invoiceService) InvoiceByID(ctx context.Context, id int) (*dto.InvoiceDetail, error) {
	return nil, nil
}

func (in *invoiceService) AddInvoice(ctx context.Context, detail *dto.InvoiceDetail) error {
	return nil
}

func (in *invoiceService) EditInvoice(ctx context.Context, detail *dto.InvoiceDetail) error {
	return nil
}
