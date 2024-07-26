package entity

import "time"

type InvoiceListsDB struct {
	ID           int       `db:"id"`
	IssueDate    time.Time `db:"issue_date"`
	Subject      string    `db:"subject"`
	TotalItems   int       `db:"total_items"`
	CustomerName string    `db:"name"`
	DueDate      time.Time `db:"due_date"`
	Status       string    `db:"status"`
}

type OrderDB struct {
	ID        int     `db:"id"`
	InvoiceID int     `db:"invoice_id"`
	ItemID    int     `db:"item"`
	Qty       int     `db:"qty"`
	Amount    float32 `db:"amount"`
}

type ItemsDB struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Type      string  `db:"type"`
	UnitPrice float32 `db:"unit_price"`
}

type Customer struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
}

type InvoiceDetailDB struct {
	ID           int       `db:"id"`
	IssueDate    time.Time `db:"issue_date"`
	Subject      string    `db:"subject"`
	CustomerName string    `db:"cust_name"`
	Address      string    `db:"address"`
	DueDate      time.Time `db:"due_date"`
	TotalItems   int       `db:"total_items"`
	SubTotal     float32   `db:"subtotal"`
	GrandTotal   float32   `db:"grand_total"`
}

type InvoiceGetDB struct {
	ID           int       `db:"id"`
	IssueDate    time.Time `db:"issue_date"`
	Subject      string    `db:"subject"`
	CustomerName string    `db:"cust_name"`
	DueDate      time.Time `db:"due_date"`
	Status       string    `db:"status"`
}
