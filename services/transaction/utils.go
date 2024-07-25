package transaction

import (
	"errors"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/loco/db"
)

func createTransaction(payload CreateTransactionReq) (code int, err error) {
	var conn *pg.DB

	trans := db.Transaction{
		TransactionID:   payload.ID,
		TransactionType: payload.Type,
		ParentID:        payload.ParentID,
		Amount:          payload.Amount,
	}

	if conn, err = db.Conn(); err != nil {
		code = http.StatusInternalServerError
		return
	}

	if _, err = conn.Model(&trans).
		OnConflict("ON CONSTRAINT transaction_pkey DO UPDATE").
		Set("amount = EXCLUDED.amount").
		Insert(); err != nil {
		code = http.StatusInternalServerError
		return
	}

	return
}

func getTransaction(id int64) (resp db.Transaction, code int, err error) {
	var conn *pg.DB

	if conn, err = db.Conn(); err != nil {
		code = http.StatusInternalServerError
		return
	}

	if err = conn.Model(&resp).Where("transaction_id = ?", id).Select(); err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			err = nil
			return
		}
		code = http.StatusInternalServerError
		return
	}

	return
}

func getTypeSum(transactionType string) (amount float64, code int, err error) {
	var conn *pg.DB

	if conn, err = db.Conn(); err != nil {
		code = http.StatusInternalServerError
		return
	}

	query := `SELECT SUM(amount) FROM transaction WHERE transaction_type = ? ;`

	if _, err = conn.Query(pg.Scan(&amount), query, transactionType); err != nil {
		code = http.StatusInternalServerError
		return
	}

	return
}

func getTransactionSum(id int64) (amount float64, code int, err error) {
	var conn *pg.DB

	if conn, err = db.Conn(); err != nil {
		code = http.StatusInternalServerError
		return
	}

	query := `WITH RECURSIVE transaction_tree AS (
				SELECT transaction_id, parent_id, amount
				FROM transaction
				WHERE transaction_id = ?
				UNION ALL
				SELECT t.transaction_id, t.parent_id, t.amount
				FROM transaction t
				JOIN transaction_tree tt ON tt.transaction_id = t.parent_id AND t.parent_id <> 0
			)
			SELECT SUM(amount) AS total_amount
			FROM transaction_tree;`

	if _, err = conn.Query(pg.Scan(&amount), query, id); err != nil {
		code = http.StatusInternalServerError
		return
	}

	return
}
