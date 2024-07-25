package transaction

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loco/lib"
)

func pCreateTransaction(req *gin.Context) (data any, code int, err error) {
	var reqPayload CreateTransactionReq
	if err = req.ShouldBindJSON(&reqPayload); err != nil {
		code = http.StatusBadRequest
		return
	}

	if reqPayload.ID, err = strconv.ParseInt(req.Param("id"), 10, 64); err != nil {
		code = http.StatusBadRequest
		return
	}

	if reqPayload.ID <= 0 {
		code = http.StatusBadRequest
		err = errors.New("invalid transaction id")
		return
	}

	if code, err = createTransaction(reqPayload); err != nil {
		return
	}

	data = lib.Success

	return
}

func pGetTransaction(req *gin.Context) (data any, code int, err error) {
	var id int64

	if id, err = strconv.ParseInt(req.Param("id"), 10, 64); err != nil {
		code = http.StatusBadRequest
		return
	}

	if id <= 0 {
		code = http.StatusBadRequest
		err = errors.New("invalid transaction id")
		return
	}

	if data, code, err = getTransaction(id); err != nil {
		return
	}

	return
}

func pGetTypeSum(req *gin.Context) (data any, code int, err error) {

	transactionType := req.Param("type")

	if transactionType == "" {
		code = http.StatusBadRequest
		err = errors.New("empty type")
		return
	}

	if data, code, err = getTypeSum(transactionType); err != nil {
		return
	}

	return
}

func pGetTransactionSum(req *gin.Context) (data any, code int, err error) {
	var id int64

	if id, err = strconv.ParseInt(req.Param("id"), 10, 64); err != nil {
		code = http.StatusBadRequest
		return
	}

	if id <= 0 {
		code = http.StatusBadRequest
		err = errors.New("invalid transaction id")
		return
	}

	if data, code, err = getTransactionSum(id); err != nil {
		return
	}

	return
}
