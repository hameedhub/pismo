package handler

import (
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
// @Param   account  body      model.Transaction  true  "Create Transaction"
// @Success 201 {object} model.Transaction
// @Router /transactions [post]
func (t transactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	if err := FromJSON(&transaction, r.Body); err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	trans, err := t.transactionService.Create(transaction)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	Response(w, http.StatusCreated, trans)
}

func NewTransactionHandler(ts service.TransactionService) TransactionHandler {
	return &transactionHandler{
		transactionService: ts,
	}
}
