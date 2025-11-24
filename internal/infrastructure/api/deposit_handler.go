package api

import (
	"bank/internal/application/usecase"
	"encoding/json"
	"net/http"
)

type DepositHandler struct {
	depositUC *usecase.DepositUseCase
}

func NewDepositHandler(depositUC *usecase.DepositUseCase) *DepositHandler {
	return &DepositHandler{depositUC: depositUC}
}

type DepositRequest struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (depositHandler *DepositHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var req DepositRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	err := depositHandler.depositUC.Execute(usecase.DepositInput{ID: req.ID, Amount: req.Amount})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
