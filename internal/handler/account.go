package handler

import (
	"errors"
	"github.com/gorilla/mux"
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
// @Param   account  body      model.Account  true  "Create Account"
// @Success 201 {object} model.Account
// @Router /accounts [post]
func (a accountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	if err := FromJSON(&account, r.Body); err != nil {
		Response(w, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	if err := account.IsValid(); err != nil {
		Response(w, http.StatusBadRequest, errors.New("document_number required").Error())
		return
	}
	create, err := a.accountService.Create(account)
	if err != nil {
		Response(w, http.StatusConflict, err.Error())
		return
	}
	Response(w, http.StatusCreated, create)
}

// GetById godoc
// @Summary Get an account by ID
// @Description Get details of an account corresponding to the input ID
// @Tags accounts
// @Produce  json
// @Param   id  path      int  true  "Account ID"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [get]
func (a accountHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	account, err := a.accountService.GetById(id)
	if err != nil {
		Response(w, http.StatusNotFound, err.Error())
		return
	}
	Response(w, http.StatusOK, account)
}

func NewAccountHandler(as service.AccountService) AccountHandler {
	return &accountHandler{accountService: as}
}
