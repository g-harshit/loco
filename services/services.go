package services

import (
	"github.com/gin-gonic/gin"
	"github.com/loco/services/transaction"
)

//InitServices : Initialise All Services
func InitServices(r *gin.Engine) {
	transactionservice := r.Group("transactionservice")
	{
		transaction.LoadServices(transactionservice)
	}
}
