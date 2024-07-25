package entity

import "time"

type InvoiceLists struct {
	ID         int       `db:"id"`
	IssueDate  time.Time `db:"issue_date"`
	Subject    string    `db:"subject"`
	TotalItems int       `db:"total_items"`
	Customer   string    `db:"customer"`
	DueDate    time.Time `db:"due_date"`
	Status     string    `db:"status"`
}

type Items struct {
	ID        int     `db:"id"`
	Item      string  `db:"item"`
	Qty       int     `db:"qty"`
	UnitPrice float32 `db:"unit_price"`
	Amount    float32 `db:"amount"`
}

type InvoiceDetail struct {
	ID int `db:"id"`
}
