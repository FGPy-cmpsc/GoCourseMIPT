package api

import (
	"bank/internal/application/usecase"
	"encoding/json"
	"net/http"
)

type TransferHandler struct {
	transferUC *usecase.TransferUseCase
}

func NewTransferHandler(transferUC *usecase.TransferUseCase) *TransferHandler {
	return &TransferHandler{transferUC: transferUC}
}

type TransferRequest struct {
	SenderID   int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
	Amount     int64 `json:"amount"`
}

func (transferHandler *TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	err := transferHandler.transferUC.Execute(usecase.TransferInput{SenderID: req.SenderID, ReceiverID: req.ReceiverID, Amount: req.Amount})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
