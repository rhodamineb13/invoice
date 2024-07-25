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
	Page       int
	Limit      int
}

type InvoiceDetailDTO struct {
	ID          int       `json:"id"`
	IssueDate   time.Time `json:"issue_date"`
	Subject     string    `json:"subject"`
	DueDate     time.Time `json:"due_date"`
	Address     string    `json:"address"`
	Summary     InvoiceSummaryDTO
	DetailItems []ItemsDTO
}

type InvoiceSummaryDTO struct {
	TotalItems int     `json:"total_item(s)"`
	SubTotal   float32 `json:"subtotal"`
	Tax        int     `json:"tax"`
	GrandTotal float32 `json:"grand_total"`
}

type ItemsDTO struct {
	Item      string  `json:"item"`
	Qty       int     `json:"qty"`
	UnitPrice float32 `json:"unit_price"`
	Amount    float32 `json:"amount"`
}
