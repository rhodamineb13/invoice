package dto

import "time"

type InvoiceListsDTO struct {
	ID         int       `json:"id"`
	IssueDate  time.Time `json:"issue_date"`
	Subject    string    `json:"subject"`
	TotalItems int       `json:"total_items"`
	Customer   string    `json:"customer"`
	DueDate    time.Time `json:"due_date"`
	Status     string    `json:"status"`
}

type InvoiceDetailDTO struct {
	ID         int       `json:"id"`
	IssueDate  time.Time `json:"issue_date"`
	Subject    string    `json:"subject"`
	DueDate    time.Time `json:"due_date"`
	Address    string    `json:"address"`
	TotalItems int       `json:"total_item(s)"`
	SubTotal   float32   `json:"subtotal"`
	Tax        int       `json:"tax"`
	GrandTotal float32   `json:"grand_total"`
	Orders     []OrdersDTO
}

type InvoiceInsertDTO struct {
	IssueDate string      `json:"issue_date"`
	Subject   string      `json:"subject"`
	ID        int         `json:"cust_id"`
	DueDate   string      `json:"due_date"`
	Address   string      `json:"address"`
	Status    string      `json:"status"`
	Orders    []OrdersDTO `json:"orders"`
}

type OrdersDTO struct {
	ItemID    int     `json:"item_id,omitempty"`
	ItemName  string  `json:"item_name,omitempty"`
	Qty       int     `json:"qty"`
	UnitPrice float32 `json:"unit_price,omitempty"`
	Amount    float32 `json:"amount,omitempty"`
}

type InvoiceUpdateDTO struct {
	IssueDate string `json:"issue_date,omitempty"`
	Subject   string `json:"subject,omitempty"`
	DueDate   string `json:"due_date,omitempty"`
	Qty       int    `json:"qty,omitempty"`
	Status    string `json:"status,omitempty"`
}
