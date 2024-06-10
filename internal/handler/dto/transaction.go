package dto

type TransactionRequest struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float32 `json:"amount"`
}

type TransactionResponse struct {
	ID              int     `json:"transaction_id"`
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float32 `json:"amount"`
}
