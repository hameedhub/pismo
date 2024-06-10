package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/hameedhub/pismo/internal/handler/dto"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/service"
	"net/http"
	"strconv"
)

type AccountHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	accountService service.AccountService
}

// Create godoc
// @Summary      Create a new account
// @Description Create a new account with the input payload
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param   account  body  dto.AccountRequest  true  "Create Account"
// @Success 201 {object} dto.AccountResponse
// @Failure 500 {object} Error
// @Failure 400 {object} Error
// @Failure 409 {object} Error
// @Router /accounts [post]
func (a accountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.AccountRequest
	if err := FromJSON(&req, r.Body); err != nil {
		ErrorResponse(w, http.StatusInternalServerError, errors.New("internal server error").Error())
		return
	}
	account := model.Account{DocumentNumber: req.DocumentNumber}
	if err := account.IsValid(); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	create, err := a.accountService.Create(account)
	if err != nil {
		ErrorResponse(w, http.StatusConflict, err.Error())
		return
	}
	Response(w, http.StatusCreated, dto.AccountResponse{ID: create.ID, DocumentNumber: create.DocumentNumber})
}

// GetById godoc
// @Summary Get an account by ID
// @Description Get details of an account corresponding to the input ID
// @Tags accounts
// @Produce  json
// @Param   id  path      int  true  "Account ID"
// @Success 200 {object} dto.AccountResponse
// @Failure 500 {object} Error
// @Failure 404 {object} Error
// @Router /accounts/{id} [get]
func (a accountHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	acc, err := a.accountService.GetById(id)
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	Response(w, http.StatusOK, dto.AccountResponse{ID: acc.ID, DocumentNumber: acc.DocumentNumber})
}

func NewAccountHandler(as service.AccountService) AccountHandler {
	return &accountHandler{accountService: as}
}
