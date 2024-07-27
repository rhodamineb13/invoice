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
	GetAllInvoices(context.Context, int, int) ([]dto.InvoiceListsDTO, error)
	SelectInvoiceByID(context.Context, int) (*dto.InvoiceDetailDTO, error)
	AddInvoice(context.Context, *dto.InvoiceInsertDTO) error
	EditInvoice(context.Context, int, *dto.InvoiceUpdateDTO) error
}

func NewInvoiceService(inv repository.InvoiceRepository) InvoiceService {
	return &invoiceService{
		inv,
	}
}

func (in *invoiceService) GetAllInvoices(ctx context.Context, page int, limit int) ([]dto.InvoiceListsDTO, error) {
	lists, err := in.invoiceRepository.GetInvoice(ctx, page, limit)
	if err == nil {
		listsDTO := []dto.InvoiceListsDTO{}
		for _, invoice := range lists {
			invDTO := dto.InvoiceListsDTO{
				ID:        invoice.ID,
				IssueDate: invoice.IssueDate,
				Subject:   invoice.Subject,
				Customer:  invoice.CustomerName,
				DueDate:   invoice.DueDate,
				Status:    invoice.Status,
			}
			listsDTO = append(listsDTO, invDTO)
		}
		return listsDTO, nil
	}

	switch err {
	case sql.ErrNoRows:
		return nil, helper.NewCustomError(http.StatusOK, "there are no invoices")
	default:
		return nil, helper.NewCustomError(http.StatusInternalServerError, err.Error())
	}
}

func (in *invoiceService) SelectInvoiceByID(ctx context.Context, id int) (*dto.InvoiceDetailDTO, error) {
	if invEntity, err := in.invoiceRepository.SearchInvoice(ctx, id); err == nil {
		invDTO := &dto.InvoiceDetailDTO{
			ID:         invEntity.ID,
			IssueDate:  invEntity.IssueDate,
			Subject:    invEntity.Subject,
			DueDate:    invEntity.DueDate,
			TotalItems: invEntity.TotalItems,
			SubTotal:   invEntity.SubTotal,
			Tax:        10,
			GrandTotal: invEntity.GrandTotal,
		}
		for _, items := range invEntity.Orders {
			item := dto.OrdersDTO{
				ItemName:  items.ItemName,
				Qty:       items.Qty,
				UnitPrice: items.UnitPrice,
				Amount:    items.Amount,
			}
			invDTO.Orders = append(invDTO.Orders, item)
		}

		return invDTO, nil
	}

	return nil, helper.NewCustomError(http.StatusNotFound, "invoice not found")
}

func (in *invoiceService) AddInvoice(ctx context.Context, detail *dto.InvoiceInsertDTO) error {
	issue, err := helper.ParseTime(detail.IssueDate)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, fmt.Sprintf("fail to parse time format %s", detail.IssueDate))
	}

	due, err := helper.ParseTime(detail.DueDate)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, fmt.Sprintf("fail to parse time format %s", detail.DueDate))
	}
	invEntity := &entity.InvoiceInsertDB{
		IssueDate:  issue,
		CustomerID: detail.ID,
		Subject:    detail.Subject,
		DueDate:    due,
		Address:    detail.Address,
		Status:     detail.Status,
		Orders:     make([]entity.OrderDB, 0),
	}
	for _, ord := range detail.Orders {
		ordDB := entity.OrderDB{
			ItemID: ord.ItemID,
			Qty:    ord.Qty,
		}

		invEntity.Orders = append(invEntity.Orders, ordDB)
	}

	if err := in.invoiceRepository.InsertInvoice(ctx, invEntity); err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("unexpected error in adding invoice: %w", err))
	}

	return nil
}

func (in *invoiceService) EditInvoice(ctx context.Context, invID int, up *dto.InvoiceUpdateDTO) error {
	issue, err := helper.ParseTime(up.IssueDate)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, fmt.Sprintf("fail to parse time format %s", up.IssueDate))
	}

	due, err := helper.ParseTime(up.DueDate)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, fmt.Sprintf("fail to parse time format %s", up.DueDate))
	}

	invEntity := &entity.InvoiceUpdateDB{
		IssueDate: issue,
		Subject:   up.Subject,
		DueDate:   due,
	}

	if err := in.invoiceRepository.UpdateInvoice(ctx, invID, invEntity); err != nil {
		switch err {
		case sql.ErrNoRows:
			return helper.NewCustomError(http.StatusBadRequest, fmt.Sprintf("can't edit the requested invoice with ID %d", invID))
		default:
			return helper.NewCustomError(http.StatusInternalServerError, "unexpected error in updating invoice")
		}
	}

	return nil
}
