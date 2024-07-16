package controller

import (
	"demo_fork/model"

	"github.com/gin-gonic/gin"
)

func GetTransactions() func(c *gin.Context) {
	return func(c *gin.Context) {

		transactions := []model.Transaction{}

		var count int64 = 0

		c.JSON(200, gin.H{
			"code":    200,
			"message": "Success",
			"data":    transactions,
			"count":   count,
		})
		return
	}
}
func GetTransaction() func(c *gin.Context) {
	return func(c *gin.Context) {

		transaction := model.Transaction{}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "Success",
			"data":    transaction,
		})
	}
}
func CreateTransaction() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(200, gin.H{
			"code":    200,
			"message": "Success",
		})
	}
}
