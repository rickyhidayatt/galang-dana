package handler

import (
	"net/http"

	"github.com/galang-dana/domain/formatter"
	"github.com/galang-dana/domain/input"
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
	response := utils.ApiResponse("List of campaigns", http.StatusOK, "success", formatter.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (ca *campaignHandler) GetCampaign(c *gin.Context) {
	var input input.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := utils.ApiResponse("error get campaign by id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)

		return
	}

	campaign, err := ca.campaignUsecase.GetCampaignById(input)
	if err != nil {

		response := utils.ApiResponse("error get campaign by id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Campaign Detail", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
}
