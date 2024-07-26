package repository

import (
	"context"
	"database/sql"
	"invoice/common/entity"
	"invoice/helper"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type invoiceRepository struct {
	db *sqlx.DB
}

type InvoiceRepository interface {
	GetInvoice(context.Context, int, int) ([]entity.InvoiceListsDB, error)
	SelectInvoice(context.Context, int) (*entity.InvoiceDetailDB, error)
	InsertInvoice(context.Context, *entity.InvoiceDetailDB) error
	UpdateInvoice(context.Context, *entity.InvoiceDetailDB) error
}

func NewInvoiceRepository(db *sqlx.DB) InvoiceRepository {
	return &invoiceRepository{
		db,
	}
}

func (in *invoiceRepository) GetInvoice(ctx context.Context, page int, limit int) ([]entity.InvoiceListsDB, error) {
	return nil, nil
}

func (in *invoiceRepository) SelectInvoice(ctx context.Context, id int) (*entity.InvoiceDetailDB, error) {
	return nil, nil
}

func (in *invoiceRepository) InsertInvoice(ctx context.Context, detail *entity.InvoiceDetailDB) error {
	queryExecute := `INSERT INTO invoice(issue_date, subject, customer_id, due_date)
	VALUES
	(?, ?, ?, ?, ?)`
	tx, err := in.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})
	if err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "unexpected error in creating transaction")
	}
	_, err = tx.ExecContext(ctx, queryExecute, detail.IssueDate, detail.Subject, detail.CustomerID, detail.DueDate)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, "bad request, fix the request and try again")
	}
	return nil
}

func (in *invoiceRepository) UpdateInvoice(ctx context.Context, detail *entity.InvoiceDetailDB) error {
	var dbInvoice entity.InvoiceDetailDB
	var items entity.ItemsDB

	querySelectIDInvoice := `SELECT * FROM invoice
	WHERE id = ? AND customer_id = ?`

	querySelectIDItem := `SELECT * FROM items
	WHERE id = ?`

	err := in.db.SelectContext(ctx, &dbInvoice, querySelectIDInvoice, detail.ID, detail.CustomerID)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, "cannot find the requested ")
	}

	err = in.db.SelectContext(ctx, &items, querySelectIDItem, detail.ID, detail.CustomerID)

	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, "cannot find the requested ")
	}

	if detail.Subject == "" {
		detail.Subject = dbInvoice.Subject
	}

	if detail.DueDate.IsZero() {
		detail.DueDate = dbInvoice.DueDate
	}

	return nil
}
