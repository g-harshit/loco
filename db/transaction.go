package db

import "time"

type Transaction struct {
	tableName       struct{}  `pg:"transaction" sql:"transaction"`
	TransactionID   int64     `json:"-" sql:"transaction_id,type:BIGINT PRIMARY KEY" pg:"transaction_id,type:BIGINT PRIMARY KEY"`
	Amount          float64   `json:"amount" sql:"amount" pg:"amount,type:numeric NOT NULL DEFAULT 0.0"`
	ParentID        int64     `json:"parent_id" sql:"parent_id" pg:"parent_id,type:int REFERENCES transaction(transaction_id) ON DELETE UPDATE CASCADE ON UPDATE CASCADE"`
	TransactionType string    `json:"transaction_type" sql:"transaction_type" pg:"transaction_type,type:varchar(100) NOT NULL"`
	CreatedAt       time.Time `json:"-" sql:"created_at" pg:"created_at,type:timestamp NOT NULL DEFAULT NOW()"`
}

//Index of the table. For composite index use ,
//Default index type is btree. For gin index use gin value
func (Transaction) Index() map[string]string {
	idx := map[string]string{}
	return idx
}
