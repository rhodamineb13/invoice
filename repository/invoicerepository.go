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
	var invList []entity.InvoiceListsDB
	offset := limit * (page - 1)

	queryGet := `SELECT in.id, in.issue_date, in.subject, it.count, c.name, in.due_date, in.status
	FROM invoice in
	INNER JOIN (SELECT invoice_id, COUNT(items.invoice_id) AS count from items GROUP BY invoice_id) it on in.id = it.invoice_id
	INNER JOIN customer c ON in.customer_id = c.id
	ORDER BY id DESC
	LIMIT $ OFFSET $;`

	err := in.db.GetContext(ctx, invList, queryGet, limit, offset)
	if err != nil {
		return nil, err
	}

	return invList, nil
}

func (in *invoiceRepository) SelectInvoice(ctx context.Context, id int) (*entity.InvoiceDetailDB, error) {
	var invDetail *entity.InvoiceDetailDB

	querySelect := `SELECT in.id, in.issue_date, in.subject, c.name AS cust_name, c.address, in.due_date, COUNT(o.ID) as total_item, SUM(o.qty*o.amount) as subtotal, subtotal*0.9 as grand_total
	FROM invoice in
	INNER JOIN (SELECT * FROM order WHERE invoice_id = $) o
	INNER JOIN customer c ON in.customer_id = c.id;`

	err := in.db.SelectContext(ctx, &invDetail, querySelect, id)
	if err != nil {
		return nil, err
	}

	return invDetail, nil
}

func (in *invoiceRepository) InsertInvoice(ctx context.Context, detail *entity.InvoiceDetailDB) error {
	queryExecute := `INSERT INTO invoice(issue_date, subject, cust_id, due_date)
	VALUES
	(?, ?, ?, ?, ?)`
	tx, err := in.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})
	if err != nil {
		return err
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
	WHERE id = ?`

	querySelectIDItem := `SELECT * FROM items
	WHERE id = ?`

	err := in.db.SelectContext(ctx, &dbInvoice, querySelectIDInvoice, detail.ID)
	if err != nil {
		return helper.NewCustomError(http.StatusBadRequest, "cannot find the requested ")
	}

	err = in.db.SelectContext(ctx, &items, querySelectIDItem, detail.ID)

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
