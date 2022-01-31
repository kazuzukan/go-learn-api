package handler

import (
	"bwa-project/helper"
	"bwa-project/transaction"
	"bwa-project/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transctionService transaction.Service
}

func NewtransctionHandler(service transaction.Service) transactionHandler {
	return transactionHandler{service}
}

func (h transactionHandler) GetCampaignTransctions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transcactions, err := h.transctionService.GetTransactionByCampaignId(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	formatter := transaction.FormatCampaignTransctions(transcactions)
	response := helper.APIResponse("Campaign transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	transactions, err := h.transctionService.GetTransactionByUserId(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	formatter := transaction.FormatUserTransctions(transactions)
	response := helper.APIResponse("Campaign transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		// fmt.Println(errorMessage)

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.transctionService.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transactions", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	formatter := transaction.FormatTransaction(newTransaction)
	response := helper.APIResponse("Success to create transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		// fmt.Println(errorMessage)

		response := helper.APIResponse("Failed to process notification", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.transctionService.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadGateway, "failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
