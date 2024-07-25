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
	GetInvoice(context.Context) ([]entity.InvoiceListsDB, error)
	SelectInvoice(context.Context, int) (*entity.InvoiceDetailDB, error)
	InsertInvoice(context.Context, *entity.InvoiceDetailDB) error
	UpdateInvoice(context.Context, *entity.InvoiceDetailDB) error
}

func NewInvoiceRepository(db *sqlx.DB) InvoiceRepository {
	return &invoiceRepository{
		db,
	}
}

func (in *invoiceRepository) GetInvoice(ctx context.Context) ([]entity.InvoiceListsDB, error) {
	return nil, nil
}

func (in *invoiceRepository) SelectInvoice(ctx context.Context, id int) (*entity.InvoiceDetailDB, error) {
	return nil, nil
}

func (in *invoiceRepository) InsertInvoice(ctx context.Context, detail *entity.InvoiceDetailDB) error {
	return nil
}

func (in *invoiceRepository) UpdateInvoice(ctx context.Context, detail *entity.InvoiceDetailDB) error {
	return nil
}
