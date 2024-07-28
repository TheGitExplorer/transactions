package routes

import (
	"net/http"
	"strconv"

	"github.com/TheGitExplorer/transactions/entity"
	"github.com/TheGitExplorer/transactions/service"
	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(router *gin.Engine) {
	router.PUT("/transactionservice/transaction/:id", putTransaction)
	router.GET("/transactionservice/transaction/:id", getTransaction)
	router.GET("/transactionservice/types/:type", getTransactionsByType)
	router.GET("/transactionservice/sum/:id", getTransactionSum)
}

func putTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}

	var transaction entity.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	transaction.ID = id

	err = service.AddTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func getTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}

	transaction, exists, err := service.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func getTransactionsByType(c *gin.Context) {
	transactionType := c.Param("type")
	ids, err := service.GetTransactionsByType(transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ids)
}

func getTransactionSum(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}

	sum, err := service.CalculateSum(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
