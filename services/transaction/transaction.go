package transaction

import (
	"github.com/gin-gonic/gin"
)

func LoadServices(r *gin.RouterGroup) {
	r.PUT("transaction/:id", CreateTransaction)
	r.GET("transaction/:id", GetTransaction)
	r.GET("types/:type", GetTransactionTypeSum)
	r.GET("sum/:id", GetTransactionSum)
}

func CreateTransaction(c *gin.Context) {
	data, code, err := pCreateTransaction(c)
	if err != nil {
		c.JSON(code, err.Error())
	} else {
		c.JSON(code, data)
	}
}

func GetTransaction(c *gin.Context) {
	data, code, err := pGetTransaction(c)

	if err != nil {
		c.JSON(code, err.Error())
	} else {
		c.JSON(code, data)
	}
}

func GetTransactionTypeSum(c *gin.Context) {
	data, code, err := pGetTypeSum(c)

	if err != nil {
		c.JSON(code, err.Error())
	} else {
		c.JSON(code, data)
	}
}

func GetTransactionSum(c *gin.Context) {
	data, code, err := pGetTransactionSum(c)

	if err != nil {
		c.JSON(code, err.Error())
	} else {
		c.JSON(code, data)
	}
}
