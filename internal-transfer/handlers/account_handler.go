package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r.paparao/internal-transfer/db"
	"github.com/r.paparao/internal-transfer/models"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO accounts (account_id, balance) VALUES ($1, $2)", acc.AccountID, acc.Balance)
	if err != nil {
		http.Error(w, "Unable to create account", http.StatusInternalServerError)
		return
	}

	// Custom response
	type AccountResponse struct {
		AccountID      int64  `json:"account_id"`
		InitialBalance string `json:"initial_balance"`
	}

	resp := AccountResponse{
		AccountID:      acc.AccountID,
		InitialBalance: fmt.Sprintf("%.5f", acc.Balance),
	}

	w.Header().Set("Content-Type", "application/json") // ✅ Add this
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}




func GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["account_id"]

	var acc models.Account
	err := db.DB.QueryRow("SELECT account_id, balance FROM accounts WHERE account_id = $1", accountID).
		Scan(&acc.AccountID, &acc.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	type AccountResponse struct {
		AccountID int64  `json:"account_id"`
		Balance   string `json:"balance"`
	}

	resp := AccountResponse{
		AccountID: acc.AccountID,
		Balance:   fmt.Sprintf("%.5f", acc.Balance),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
