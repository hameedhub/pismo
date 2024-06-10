package handler

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/hameedhub/pismo/internal/config"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/repository"
	"github.com/hameedhub/pismo/internal/repository/database"
	"github.com/hameedhub/pismo/internal/service"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupDatabase() (*gorm.DB, error) {
	cfg := config.ReadConfig("../../")
	return database.Run(cfg), nil
}
func setupTestEnvironment(t *testing.T) (*gorm.DB, func()) {
	t.Helper()

	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	return db, func() {
		sqlDB, err := db.DB()
		if err != nil {
			t.Fatalf("failed to get db connection: %v", err)
		}
		db.Migrator().DropTable(&model.Account{})
		sqlDB.Close()
	}
}
func createRequestAndRecorder(body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	return req, resp
}

const body = `{"document_number":"12332231"}`
const invalidBody = `{"document_number":""}`

func TestAccountHandler_Create_Invalid(t *testing.T) {
	db, cleanup := setupTestEnvironment(t)
	defer cleanup()
	accountService := service.NewAccountService(repository.Init(db))
	handler := NewAccountHandler(accountService)

	req, resp := createRequestAndRecorder(invalidBody)
	handler.Create(resp, req)
	if resp.Code != http.StatusBadRequest {
		t.Errorf("expected status %v; got %v", http.StatusBadRequest, resp.Code)
	}
}
func TestAccountHandler_Create(t *testing.T) {
	db, cleanup := setupTestEnvironment(t)
	defer cleanup()
	accountService := service.NewAccountService(repository.Init(db))
	handler := NewAccountHandler(accountService)

	req, resp := createRequestAndRecorder(body)
	handler.Create(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, resp.Code)
	}
}

func TestAccountHandler_Create_Existing(t *testing.T) {
	db, cleanup := setupTestEnvironment(t)
	defer cleanup()
	accountService := service.NewAccountService(repository.Init(db))
	handler := NewAccountHandler(accountService)

	req, resp := createRequestAndRecorder(body)
	handler.Create(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, resp.Code)
	}
	req, resp = createRequestAndRecorder(body)
	handler.Create(resp, req)
	if resp.Code != http.StatusConflict {
		t.Errorf("expected status %v; got %v", http.StatusConflict, resp.Code)
	}
}
func TestAccountHandler_GetById_NotFound(t *testing.T) {
	db, cleanup := setupTestEnvironment(t)
	defer cleanup()
	accountService := service.NewAccountService(repository.Init(db))
	handler := NewAccountHandler(accountService)
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", handler.GetById).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil) // GET requests usually have no body.
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("expected status %v; got %v", http.StatusNotFound, resp.Code)
	}
}

func TestAccountHandler_GetById(t *testing.T) {
	db, cleanup := setupTestEnvironment(t)
	defer cleanup()
	accountService := service.NewAccountService(repository.Init(db))
	handler := NewAccountHandler(accountService)

	req, resp := createRequestAndRecorder(body)
	handler.Create(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, resp.Code)
	}
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", handler.GetById).Methods(http.MethodGet)

	req = httptest.NewRequest(http.MethodGet, "/accounts/1", nil) // GET requests usually have no body.
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status %v; got %v", http.StatusOK, resp.Code)
	}

}
