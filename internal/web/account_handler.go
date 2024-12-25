package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rafael-italiano/brynhildr/internal/service"
)

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.service.GetAccounts()
	if err != nil {
		http.Error(w, "failed to get accounts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account service.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateAccount(&account)
	if err != nil {
		http.Error(w, "failed to create account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := h.service.GetAccountByID(id)
	if err != nil {
		http.Error(w, "failed to get account", http.StatusInternalServerError)
		return
	}
	if account == nil {
		http.Error(w, "account not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid account ID", http.StatusBadRequest)
		return
	}

	var account service.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
	}
	account.ID = id

	if err := h.service.UpdateAccount(&account); err != nil {
		http.Error(w, "failed to upddate account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid account ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteAccount(id); err != nil {
		http.Error(w, "failed to delete account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
