package api

import (
	"bank/internal/application/usecase"
	"encoding/json"
	"net/http"
)

type BalanceHandler struct {
	balanceUC *usecase.GetBalanceUseCase
}

func NewBalanceHandler(balanceUC *usecase.GetBalanceUseCase) *BalanceHandler {
	return &BalanceHandler{balanceUC: balanceUC}
}

type GetBalanceRequest struct {
	ID int64 `json:"id"`
}

type GetBalanceResponse struct {
	Balance int64 `json:"balance"`
}

func (balanceHandler *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var req GetBalanceRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	out, err := balanceHandler.balanceUC.Execute(usecase.GetBalanceInput{ID: req.ID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := GetBalanceResponse{Balance: out.Balance}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
