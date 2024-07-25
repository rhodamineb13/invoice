package service

import (
	"context"
	"database/sql"
	"fmt"
	"invoice/common/dto"
	"invoice/common/entity"
	"invoice/helper"
	"invoice/repository"
	"net/http"
)

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

type InvoiceService interface {
	InvoiceIndex(context.Context) ([]dto.InvoiceListsDTO, error)
	InvoiceByID(context.Context, int) (*dto.InvoiceDetailDTO, error)
	AddInvoice(context.Context, *dto.InvoiceDetailDTO) error
	EditInvoice(context.Context, *dto.InvoiceDetailDTO) error
}

func NewInvoiceService(inv repository.InvoiceRepository) InvoiceService {
	return &invoiceService{
		inv,
	}
}

func (in *invoiceService) InvoiceIndex(ctx context.Context) ([]dto.InvoiceListsDTO, error) {
	lists, err := in.invoiceRepository.GetInvoice(ctx)
	if err != nil {
		listsDTO := []dto.InvoiceListsDTO{}
		for _, invoice := range lists {
			invDTO := dto.InvoiceListsDTO{
				ID:         invoice.ID,
				IssueDate:  invoice.IssueDate,
				Subject:    invoice.Subject,
				TotalItems: invoice.TotalItems,
				Customer:   invoice.Customer,
				DueDate:    invoice.DueDate,
				Status:     invoice.Status,
			}
			listsDTO = append(listsDTO, invDTO)
		}
		return listsDTO, nil
	}

	switch err {
	case sql.ErrNoRows:
		return nil, helper.NewCustomError(http.StatusOK, "there are no invoices")
	default:
		return nil, helper.NewCustomError(http.StatusInternalServerError, "unexpected error")
	}
}

func (in *invoiceService) InvoiceByID(ctx context.Context, id int) (*dto.InvoiceDetailDTO, error) {
	if invEntity, err := in.invoiceRepository.SelectInvoice(ctx, id); err == nil {
		invDTO := &dto.InvoiceDetailDTO{
			ID:        invEntity.ID,
			IssueDate: invEntity.IssueDate,
			Subject:   invEntity.Subject,
			DueDate:   invEntity.DueDate,
			Address:   invEntity.Address,
		}
		return invDTO, nil
	}

	return nil, helper.NewCustomError(http.StatusNotFound, "invoice not found")
}

func (in *invoiceService) AddInvoice(ctx context.Context, detail *dto.InvoiceDetailDTO) error {
	invEntity := &entity.InvoiceDetailDB{
		IssueDate: detail.IssueDate,
		Subject:   detail.Subject,
		DueDate:   detail.DueDate,
		Address:   detail.Address,
	}
	if err := in.invoiceRepository.InsertInvoice(ctx, invEntity); err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "unexpected error in adding invoice")
	}

	return nil
}

func (in *invoiceService) EditInvoice(ctx context.Context, detail *dto.InvoiceDetailDTO) error {
	invEntity := &entity.InvoiceDetailDB{
		ID:        detail.ID,
		IssueDate: detail.IssueDate,
		Subject:   detail.Subject,
		DueDate:   detail.DueDate,
		Address:   detail.Address,
	}

	if err := in.invoiceRepository.UpdateInvoice(ctx, invEntity); err != nil {
		switch err {
		case sql.ErrNoRows:
			return helper.NewCustomError(http.StatusBadRequest, fmt.Sprintf("can't edit the requested invoice with ID %d", detail.ID))
		default:
			return helper.NewCustomError(http.StatusInternalServerError, "unexpected error in updating invoice")
		}
	}

	return nil
}
