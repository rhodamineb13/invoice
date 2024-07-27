package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"invoice/common/entity"
	"invoice/helper"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type invoiceRepository struct {
	db *sqlx.DB
}

type InvoiceRepository interface {
	GetInvoice(context.Context, int, int) ([]entity.InvoiceGetDB, error)
	SearchInvoice(context.Context, int) (*entity.InvoiceDetailDB, error)
	InsertInvoice(context.Context, *entity.InvoiceInsertDB) error
	UpdateInvoice(context.Context, int, *entity.InvoiceUpdateDB) error
}

func NewInvoiceRepository(db *sqlx.DB) InvoiceRepository {
	return &invoiceRepository{
		db,
	}
}

func (in *invoiceRepository) GetInvoice(ctx context.Context, page int, limit int) ([]entity.InvoiceGetDB, error) {
	var invList []entity.InvoiceGetDB
	offset := limit * (page - 1)

	queryGet := `SELECT invoices.id, invoices.issue_date, invoices.subject, o.total_items, c.name as cust_name, invoices.due_date, invoices.status
	FROM invoices
	INNER JOIN 
		(SELECT invoice_id, COUNT(orders.invoice_id) AS total_items from orders GROUP BY invoice_id) AS o on invoices.id = o.invoice_id
	INNER JOIN 
		customers AS c ON invoices.cust_id = c.id
	ORDER BY id DESC
	LIMIT ? OFFSET ?;`

	err := in.db.SelectContext(ctx, &invList, queryGet, limit, offset)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return invList, nil
}

func (in *invoiceRepository) SearchInvoice(ctx context.Context, id int) (*entity.InvoiceDetailDB, error) {
	var invDetail *entity.InvoiceDetailDB

	querySelect := `SELECT in.id, in.cust_id, in.issue_date, in.subject, c.name AS cust_name, c.address, in.due_date, COUNT(o.ID) as total_items, SUM(o.qty*o.amount) as subtotal, SUM(o.qty*o.amount)*0.9 as grand_total
	FROM invoice in
	INNER JOIN (SELECT * FROM order WHERE invoice_id = $) o
	INNER JOIN customer c ON in.cust_id = c.id;`

	err := in.db.SelectContext(ctx, &invDetail, querySelect, id)
	if err != nil {
		return nil, err
	}

	return invDetail, nil
}

func (in *invoiceRepository) InsertInvoice(ctx context.Context, ins *entity.InvoiceInsertDB) error {

	queryInsertInvoice := `INSERT INTO invoices(issue_date, subject, cust_id, due_date, status)
	VALUES
	(?, ?, ?, ?, 'unpaid');`

	tx, err := in.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})

	if err != nil {
		return err
	}

	_, err1 := tx.ExecContext(ctx, queryInsertInvoice, ins.IssueDate, ins.Subject, ins.CustomerID, ins.DueDate)
	if err1 == nil {
		tx.Commit()
		var invID int
		querySelectInvoice := `SELECT id FROM invoices WHERE cust_id = ?`

		if err := in.db.SelectContext(ctx, &invID, querySelectInvoice, ins.CustomerID); err != nil {
			return helper.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("unexpected error: customer id not found"))
		}
		rows, err2 := tx.QueryContext(ctx, `SELECT id FROM invoices WHERE issue_date = ? AND subject = ? AND cust_id = ? AND due_date = ? AND status = 'unpaid'`,
			ins.IssueDate, ins.Subject, ins.CustomerID, ins.DueDate)
		if err != nil {
			tx.Rollback()
			return err2
		}

		queryNewOrder := `INSERT INTO ORDER(invoice_id, item_id, qty)
		VALUES
		(?, ?, ?)`

		rows.Scan(&invID)
		for _, ord := range ins.Orders {
			_, err := tx.ExecContext(ctx, queryNewOrder, invID, ord.ItemID, ord.Qty)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	}

	return err1
}

func (in *invoiceRepository) UpdateInvoice(ctx context.Context, invID int, update *entity.InvoiceUpdateDB) error {
	detail, err := in.SearchInvoice(ctx, invID)
	if err != nil {
		return err
	}

	if update.IssueDate.IsZero() {
		update.IssueDate = detail.IssueDate
	}

	if update.DueDate.IsZero() {
		update.DueDate = detail.DueDate
	}

	if update.Subject == "" {
		update.Subject = detail.Subject
	}

	if update.CustomerID == 0 {
		update.CustomerID = detail.CustomerID
	}

	queryUpdateInvoice := `UPDATE invoices
	SET issue_date = ?, subject = ?, cust_id = ?, due_date = ?
	WHERE id = ?`

	tx, err := in.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryUpdateInvoice, update.IssueDate, update.Subject, update.CustomerID, update.DueDate, invID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (in *invoiceRepository) checkCustomerIfExists(ctx context.Context, name string, address string) (int, error) {
	var custID int
	queryCheckCustomer := `SELECT id FROM customers WHERE name = ? AND address = ?`

	err := in.db.SelectContext(ctx, &custID, queryCheckCustomer, name, address)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		queryAddCustomer := `INSERT INTO customers(name, address)
		VALUES
		(?, ?);
		SELECT LAST_INSERT_ID();`

		tx, err := in.db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.IsolationLevel(4),
		})
		if err != nil {
			return -1, err
		}

		res, err := tx.ExecContext(ctx, queryAddCustomer, name, address)
		if err != nil {
			return -1, err
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return -1, err
		}
		custID = int(lastID)
		return custID, nil
	}

	return custID, nil
}
