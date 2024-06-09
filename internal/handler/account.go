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
