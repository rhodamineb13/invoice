package repository

import (
	"context"
	"database/sql"
	"invoice/common/entity"
	"invoice/helper"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type orderRepository struct {
	db *sqlx.DB
}

type OrderRepository interface {
	GetOrders(context.Context, int) ([]entity.OrderDB, error)
	Insert(context.Context, int, *entity.OrderDB) error
	Update(context.Context, int, *entity.InvoiceOrderUpdateDB) error
	Delete(context.Context, int) error
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db,
	}
}

func (o *orderRepository) GetOrders(ctx context.Context, invID int) ([]entity.OrderDB, error) {
	var ord []entity.OrderDB
	querySelect := `SELECT o.id, i.name, o.qty, i.unit_price, o.qty*i.unit_price AS amount FROM orders AS O 
	INNER JOIN items AS i ON o.item_id = i.id
	WHERE invoice_id = ?;`

	err := o.db.SelectContext(ctx, &ord, querySelect, invID)
	if err != nil {
		return nil, err
	}

	return ord, nil
}

func (o *orderRepository) Insert(ctx context.Context, invID int, ord *entity.OrderDB) error {
	queryInsert := `INSERT INTO orders(invoice_id, item_id, qty)
	VALUES
	(?, ?, ?);`

	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryInsert, invID, ord.ItemID, ord.Qty)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) Update(ctx context.Context, ordID int, update *entity.InvoiceOrderUpdateDB) error {
	var ord *entity.InvoiceOrderUpdateDB

	querySearchOrder := `SELECT item_id, qty FROM orders
	WHERE id = ?`

	if err := o.db.GetContext(ctx, &ord, querySearchOrder, &ordID); err != nil {
		return err
	}

	if update.ItemID == 0 {
		update.ItemID = ord.ItemID
	}

	if update.Qty == 0 {
		update.Qty = ord.Qty
	}

	queryUpdateOrder := `UPDATE orders
	SET item_id = ?, qty = ?
	WHERE id = ?`

	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, queryUpdateOrder, update.ItemID, update.Qty)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (o *orderRepository) Delete(ctx context.Context, ordID int) error {
	queryDelete := `DELETE FROM orders WHERE id = ?`

	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(4),
	})

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryDelete, ordID)
	if err != nil {
		tx.Rollback()
		return helper.NewCustomError(http.StatusBadRequest, "invalid order to delete")
	}
	tx.Commit()
	return nil
}
