package handler

import (
	"errors"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/service"
	"net/http"
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
	err := FromJSON(&account, r.Body)
	if err := account.IsValid(); err != nil {
		Response(w, http.StatusBadRequest, errors.New("document_number required"))
	}

	if err != nil {
		Response(w, http.StatusInternalServerError, errors.New("internal server error"))
	}
	create, err := a.accountService.Create(account)
	if err != nil {
		Response(w, http.StatusConflict, err.Error())
		return
	}
	Response(w, http.StatusCreated, create)
}

func (a accountHandler) GetById(w http.ResponseWriter, r *http.Request) {

}

func NewAccountHandler(as service.AccountService) AccountHandler {
	return &accountHandler{accountService: as}
}
