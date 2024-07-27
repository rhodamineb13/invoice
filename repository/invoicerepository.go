package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"invoice/common/entity"

	"github.com/jmoiron/sqlx"
)

type invoiceRepository struct {
	db *sqlx.DB
}

type InvoiceRepository interface {
	GetInvoice(context.Context, int, int) ([]entity.InvoiceGetDB, error)
	SearchInvoice(context.Context, int) (*entity.InvoiceDetailDB, error)
	InsertInvoice(context.Context, *entity.InvoiceInsertDB) error
	UpdateInvoice(context.Context, *entity.InvoiceDetailDB) error
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

	querySelect := `SELECT in.id, in.issue_date, in.subject, c.name AS cust_name, c.address, in.due_date, COUNT(o.ID) as total_items, SUM(o.qty*o.amount) as subtotal, SUM(o.qty*o.amount)*0.9 as grand_total
	FROM invoice in
	INNER JOIN (SELECT * FROM order WHERE invoice_id = $) o
	INNER JOIN customer c ON in.customer_id = c.id;`

	err := in.db.SelectContext(ctx, &invDetail, querySelect, id)
	if err != nil {
		return nil, err
	}

	queryGetItems := `SELECT o.item_name, it.unit_price AS unit_price, o.qty, (it.unit_price * o.qty) AS amount
	FROM orders o
	INNER JOIN items it ON it.id = o.item_id`

	if err := in.db.GetContext(ctx, &invDetail.Orders, queryGetItems); err != nil {
		return nil, err
	}

	return invDetail, nil
}

func (in *invoiceRepository) InsertInvoice(ctx context.Context, ins *entity.InvoiceInsertDB) error {
	custID, err := in.checkCustomerIfExists(ctx, ins.CustomerName, ins.Address)
	if err != nil {
		return err
	}

	ins.CustomerID = custID

	queryInsertInvoice := `INSERT INTO invoices(issue_date, subject, cust_id, due_date, status)
	VALUES
	(?, ?, ?, ?, ?);
	SELECT LAST_INSERT_ID();`

	queryInsertOrders := `INSERT INTO orders(invoice_id, item_id, qty)
	VALUES
	(?, ?, ?)`

	tx, err := in.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})

	if err != nil {
		return err
	}

	res, err1 := tx.ExecContext(ctx, queryInsertInvoice, ins.IssueDate, ins.Subject, ins.CustomerID, ins.DueDate, ins.Status)
	var err2 error
	for _, order := range ins.Orders {
		invID, _ := res.LastInsertId()
		var itemID int
		querySelectItems := `SELECT id FROM items WHERE name = ?`
		err := in.db.SelectContext(ctx, &itemID, querySelectItems, order.ItemName)
		if err != nil {
			err2 = err
			break
		}

		_, err2 = tx.ExecContext(ctx, queryInsertOrders, invID, itemID, order.Qty)
	}
	if err1 != nil || err2 != nil {
		tx.Rollback()
		if err1 != nil {
			return err1
		}
		return err2
	}
	tx.Commit()

	return nil
}

func (in *invoiceRepository) UpdateInvoice(ctx context.Context, update *entity.InvoiceDetailDB) error {
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
