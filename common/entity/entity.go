package entity

import "time"

type InvoiceListsDB struct {
	ID         int       `db:"id"`
	IssueDate  time.Time `db:"issue_date"`
	Subject    string    `db:"subject"`
	TotalItems int       `db:"total_items"`
	Customer   string    `db:"customer"`
	DueDate    time.Time `db:"due_date"`
	Status     string    `db:"status"`
}

type ItemsDB struct {
	ID        int     `db:"id"`
	InvoiceID int     `db:"invoice_id"`
	Item      string  `db:"item"`
	Qty       int     `db:"qty"`
	UnitPrice float32 `db:"unit_price"`
	Amount    float32 `db:"amount"`
}

type Customer struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
}

type InvoiceSummaryDB struct {
	TotalItems int     `db:"total_item"`
	SubTotal   float32 `db:"subtotal"`
	Tax        int     `db:"tax"`
	GrandTotal float32 `db:"grand_total"`
}

type InvoiceDetailDB struct {
	ID         int       `db:"id"`
	IssueDate  time.Time `db:"issue_date"`
	Subject    string    `db:"subject"`
	CustomerID int       `db:"customer_id"`
	DueDate    time.Time `db:"due_date"`
}
