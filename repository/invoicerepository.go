package repository

import (
	"context"
	"database/sql"
	"fmt"
	"invoice/common/entity"

	"github.com/jmoiron/sqlx"
)

type invoiceRepository struct {
	db *sqlx.DB
}

type InvoiceRepository interface {
	GetInvoice(context.Context, int, int) ([]entity.InvoiceGetDB, error)
	SearchInvoice(context.Context, int) (entity.InvoiceDetailDB, error)
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

func (in *invoiceRepository) SearchInvoice(ctx context.Context, id int) (entity.InvoiceDetailDB, error) {
	var invDetail entity.InvoiceDetailDB

	querySelect := `SELECT invoices.id, invoices.cust_id, invoices.issue_date, invoices.subject, c.name AS cust_name, c.address AS address, invoices.due_date, COUNT(o.invoice_id) AS total_items, SUM(o.amount) AS subtotal, SUM(o.amount)*0.9 AS grand_total
					FROM invoices
					INNER JOIN (SELECT qty, invoice_id, item_id, it.unit_price, qty*it.unit_price AS amount FROM orders 
					INNER JOIN items AS it ON orders.item_id = it.id WHERE orders.invoice_id = ?) AS o on o.invoice_id = invoices.id
					INNER JOIN customers AS c on c.id = invoices.cust_id
					GROUP BY id;`

	err := in.db.GetContext(ctx, &invDetail, querySelect, id)
	fmt.Println(invDetail)
	if err != nil {
		fmt.Println(err)
		return invDetail, err
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

	res, err1 := tx.ExecContext(ctx, queryInsertInvoice, ins.IssueDate, ins.Subject, ins.CustomerID, ins.DueDate)
	if err1 != nil {
		return err1
	}

	invID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	queryInsertOrder := `INSERT INTO orders(invoice_id, item_id, qty)
	VALUES
	(?, ?, ?)`

	for _, ord := range ins.Orders {
		_, err := tx.ExecContext(ctx, queryInsertOrder, int(invID), ord.ItemID, ord.Qty)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
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
