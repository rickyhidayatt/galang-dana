package handler

import (
	"net/http"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/utils"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionUsecase usecase.TransactionUseCase
}

func TransactionHandler(transactionUC usecase.TransactionUseCase) *transactionHandler {
	return &transactionHandler{transactionUC}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input input.GetCampaignTransaction

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.ApiResponse("error get campaign transactio with id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transaction, err := h.transactionUsecase.GetTransactionByID(input)
	if err != nil {
		response := utils.ApiResponse("error get campaign transaction id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Transaction Detail", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)
}
