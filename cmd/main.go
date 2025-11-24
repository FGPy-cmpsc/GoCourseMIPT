package main

import (
	"bank/internal/application/usecase"
	"bank/internal/infrastructure/api"
	"bank/internal/infrastructure/db"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("./cmd/database.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	defer file.Close()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(file)

	mux := http.NewServeMux()

	fileDB := db.NewFileDB(file)

	getBalanceUC := usecase.NewGetBalanceUseCase(fileDB)
	getBalanceHandler := api.NewBalanceHandler(getBalanceUC)
	mux.HandleFunc("/api/get_balance/", getBalanceHandler.GetBalance)

	transferUC := usecase.NewTransferUseCase(fileDB)
	transferHandler := api.NewTransferHandler(transferUC)
	mux.HandleFunc("/api/transfer/", transferHandler.Transfer)

	depositUC := usecase.NewDepositUseCase(fileDB)
	depositHandler := api.NewDepositHandler(depositUC)
	mux.HandleFunc("/api/deposit/", depositHandler.Deposit)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	srv.ListenAndServe()
}
