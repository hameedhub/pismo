package handler

import (
	"github.com/hameedhub/pismo/internal/repository"
	"github.com/hameedhub/pismo/internal/service"
	"net/http"
	"testing"
)

const transaction = `{"operation_type_id":1, "account_id":1}`

func TestTransactionHandler_Create(t *testing.T) {
	db, cleanup := setupTestEnvironment(t)
	defer cleanup()

	accountService := service.NewAccountService(repository.Init(db))
	handlerAccount := NewAccountHandler(accountService)

	transactionService := service.NewTransactionService(repository.Init(db))
	handler := NewTransactionHandler(transactionService)

	req, resp := createRequestAndRecorder(body)
	handlerAccount.Create(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, resp.Code)
	}
	
	req, resp = createRequestAndRecorder(transaction)
	handler.Create(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, resp.Code)
	}

}
