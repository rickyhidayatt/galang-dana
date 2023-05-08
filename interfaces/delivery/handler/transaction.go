package handler

import (
	"net/http"

	"github.com/galang-dana/domain/formatter"
	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/utils"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionUsecase usecase.TransactionUseCase
	paymentUsecase     usecase.PaymentUseCase
}

func TransactionHandler(transactionUC usecase.TransactionUseCase, paymentUC usecase.PaymentUseCase) *transactionHandler {
	return &transactionHandler{transactionUC, paymentUC}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input input.GetCampaignTransaction

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.ApiResponse("error get campaign transactio with id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	input.User = currentUser

	transaction, err := h.transactionUsecase.GetTransactionByID(input)
	if err != nil {
		response := utils.ApiResponse("error get campaign transaction id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Transaction Detail", http.StatusOK, "success", formatter.FormatCampaignTransactions(transaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(model.User)
	userID := currentUser.Id

	transaction, err := h.transactionUsecase.GetTransactionByUserID(userID)
	if err != nil {
		response := utils.ApiResponse("failed to get user transaction", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Transaction Detail", http.StatusOK, "success", formatter.FormatUserTransactions(transaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var inputTransaction input.CreateTransactionInput
	err := c.ShouldBindJSON(&inputTransaction)
	if err != nil {
		response := utils.ApiResponse("failed to create transaction", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)
	inputTransaction.User = currentUser

	newTransaction, err := h.transactionUsecase.CreateTransaction(inputTransaction)
	if err != nil {
		response := utils.ApiResponse("failed to create transaction", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := utils.ApiResponse("transaction success", http.StatusOK, "success", formatter.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input input.TransactionNotificationInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := utils.ApiResponse("failed to process notification", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = h.paymentUsecase.ProcessPayment(input)
	if err != nil {
		response := utils.ApiResponse("failed to process notification", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, input)
}
