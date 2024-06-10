package handler

import (
	"github.com/hameedhub/pismo/internal/handler/dto"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/service"
	"net/http"
)

type TransactionHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct {
	transactionService service.TransactionService
}

// Create godoc
// @Summary      Create a new transaction
// @Description Create a new transaction with the input payload
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   account  body      dto.TransactionRequest  true  "Create Transaction"
// @Success 201 {object} dto.TransactionResponse
// @Router /transactions [post]
func (t transactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.TransactionRequest
	if err := FromJSON(&req, r.Body); err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	trans, err := t.transactionService.Create(model.Transaction{AccountId: req.AccountId, OperationTypeId: req.OperationTypeId, Amount: req.Amount})
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	Response(w, http.StatusCreated, dto.TransactionResponse{ID: trans.ID, AccountId: trans.AccountId, OperationTypeId: trans.OperationTypeId, Amount: trans.Amount})
}

func NewTransactionHandler(ts service.TransactionService) TransactionHandler {
	return &transactionHandler{
		transactionService: ts,
	}
}
