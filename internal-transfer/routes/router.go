package routes

import (
	"github.com/gorilla/mux"
	"github.com/r.paparao/internal-transfer/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/accounts", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/accounts/{account_id}", handlers.GetAccount).Methods("GET")
	router.HandleFunc("/api/transfer", handlers.TransferFundsHandler).Methods("POST")
	return router
}
