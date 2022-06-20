package app

import (
	"ashishi-banking/dto"
	"ashishi-banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = customerId

	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}

	writeResponse(w, http.StatusCreated, account)
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = customerId
	request.AccountId = accountId

	account, appError := h.service.MakeTransaction(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage)
	}

	writeResponse(w, http.StatusOK, account)
}
