package services

import (
	"errors"
	"log"

	"github.com/r.paparao/internal-transfer/db"
)



func TransferFunds(fromID, toID int64, amount float64) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE", fromID).Scan(&fromBalance)
	if err != nil {
		tx.Rollback()
		return err
	}

	if fromBalance < amount {
		tx.Rollback()
		return errors.New("insufficient funds")
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE account_id = $2", amount, fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2", amount, toID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions (source_account_id, destination_account_id, amount)
		VALUES ($1, $2, $3)`, fromID, toID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	log.Printf("Transferred %.2f from account %d to %d", amount, fromID, toID)
	return nil
}

