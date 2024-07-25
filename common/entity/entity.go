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
	Item      string  `db:"item"`
	Qty       int     `db:"qty"`
	UnitPrice float32 `db:"unit_price"`
	Amount    float32 `db:"amount"`
}

type InvoiceSummaryDB struct {
	TotalItems int     `json:"total_item(s)"`
	SubTotal   float32 `json:"subtotal"`
	Tax        int     `json:"tax"`
	GrandTotal float32 `json:"grand_total"`
}

type InvoiceDetailDB struct {
	ID         int       `db:"id"`
	IssueDate  time.Time `db:"issue_date"`
	Subject    string    `db:"subject"`
	TotalItems int       `db:"total_items"`
	Customer   string    `db:"customer"`
	DueDate    time.Time `db:"due_date"`
	Address    string    `db:"address"`
	Summary    InvoiceSummaryDB
	Items      []ItemsDB
}
