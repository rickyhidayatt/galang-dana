package handler

import (
	"net/http"

	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/utils"
	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignUsecase usecase.CampaignUseCase
}

func CampaignHandler(campaignUC usecase.CampaignUseCase) *campaignHandler {
	return &campaignHandler{campaignUC}
}

func (ca *campaignHandler) GetCampaigns(c *gin.Context) {
	userId := c.Query("user_id")

	campaigns, err := ca.campaignUsecase.GetCampaigns(userId)
	if err != nil {
		response := utils.ApiResponse("error get campaigns", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.ApiResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)

}
