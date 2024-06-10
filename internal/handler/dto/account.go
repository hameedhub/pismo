package dto

type AccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type AccountResponse struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
