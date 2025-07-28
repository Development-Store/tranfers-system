package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/r.paparao/internal-transfer/models"
	"github.com/r.paparao/internal-transfer/services"
)

func TransferFundsHandler(w http.ResponseWriter, r *http.Request) {
	var req models.TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := services.TransferFunds(req.FromAccountID, req.ToAccountID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}