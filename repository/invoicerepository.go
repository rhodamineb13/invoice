package repository

import (
	"context"
	"invoice/common/entity"

	"github.com/jmoiron/sqlx"
)

type invoiceRepository struct {
	db *sqlx.DB
}

type InvoiceRepository interface {
	GetInvoice(context.Context) ([]entity.InvoiceLists, error)
	SelectInvoice(context.Context, int) (*entity.InvoiceDetail, error)
	InsertInvoice(context.Context, *entity.InvoiceDetail) error
	UpdateInvoice(context.Context, *entity.InvoiceDetail) error
}

func NewInvoiceRepository(db *sqlx.DB) InvoiceRepository {
	return &invoiceRepository{
		db,
	}
}

func (in *invoiceRepository) GetInvoice(ctx context.Context) ([]entity.InvoiceLists, error) {
	return nil, nil
}

func (in *invoiceRepository) SelectInvoice(ctx context.Context, id int) (*entity.InvoiceDetail, error) {
	return nil, nil
}

func (in *invoiceRepository) InsertInvoice(ctx context.Context, detail *entity.InvoiceDetail) error {
	return nil
}

func (in *invoiceRepository) UpdateInvoice(ctx context.Context, detail *entity.InvoiceDetail) error {
	return nil
}
