package transaction

type CreateTransactionReq struct {
	ID       int64   `json:"transaction_id"`
	Type     string  `json:"type" binding:"required"`
	ParentID int64   `json:"parent_id"`
	Amount   float64 `json:"amount" binding:"required"`
}
