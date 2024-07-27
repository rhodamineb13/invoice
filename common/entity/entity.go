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
	ItemID    int     `db:"item_id"`
	ItemName  string  `db:"name"`
	Qty       int     `db:"qty"`
	UnitPrice float32 `db:"unit_price"`
	Amount    float32 `db:"amount"`
}

type InvoiceDetailDB struct {
	ID           int       `db:"id"`
	IssueDate    time.Time `db:"issue_date"`
	Subject      string    `db:"subject"`
	CustomerID   int       `db:"cust_id"`
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
	TotalItems   int       `db:"total_items"`
	CustomerName string    `db:"cust_name"`
	DueDate      time.Time `db:"due_date"`
	Status       string    `db:"status"`
}

type InvoiceInsertDB struct {
	IssueDate    time.Time `db:"issue_date"`
	Subject      string    `db:"subject"`
	CustomerName string    `db:"cust_name"`
	CustomerID   int       `db:"cust_id"`
	Address      string    `db:"address"`
	DueDate      time.Time `db:"due_date"`
	Status       string    `db:"status"`
	Orders       []OrderDB
}

type InvoiceOrderUpdateDB struct {
	OrderID int `db:"id"`
	ItemID  int `db:"item_id"`
	Qty     int `db:"qty"`
}

type InvoiceUpdateDB struct {
	IssueDate  time.Time `db:"issue_date"`
	DueDate    time.Time `db:"due_date"`
	Subject    string    `db:"subject"`
	CustomerID int       `db:"cust_id"`
}
