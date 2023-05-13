package handler

import (
	"fmt"
	"net/http"

	"github.com/galang-dana/domain/formatter"
	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
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

func (ca *campaignHandler) GetCampaignById(c *gin.Context) {
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

	response := utils.ApiResponse("Campaign Detail", http.StatusOK, "success", formatter.FormatCampaignDetail(campaign))
	c.JSON(http.StatusOK, response)
}

func (ca *campaignHandler) CreateCampaign(c *gin.Context) {
	var input input.CreateCampaign
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidatorError(err)
		errorsMessage := gin.H{"error": errors}
		response := utils.ApiResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)
	input.User = currentUser

	newCampaign, err := ca.campaignUsecase.CreateCampaign(input)
	if err != nil {
		response := utils.ApiResponse("failed to create campaign", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.ApiResponse("success to create campaign", http.StatusOK, "success", formatter.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (ca *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID input.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := utils.ApiResponse("failed to get update campaign", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData input.CreateCampaign
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := utils.FormatValidatorError(err)
		errorsMessage := gin.H{"error": errors}
		response := utils.ApiResponse("failed to update campaign", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	inputData.User = currentUser

	updateCampaign, err := ca.campaignUsecase.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := utils.ApiResponse("failed to update campaign", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("success to create campaign", http.StatusOK, "success", formatter.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

func (ca *campaignHandler) UploadImage(c *gin.Context) {
	var input input.CreateCampaignImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := utils.FormatValidatorError(err)
		errorsMessage := gin.H{"error": errors}
		response := utils.ApiResponse("failed to upload image", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	input.User = currentUser
	userID := currentUser.Id

	file, err := c.FormFile("file")
	if err != nil {
		errorsMessage := gin.H{"is_uploaded": false}
		response := utils.ApiResponse("failed to upload campaign image", http.StatusBadRequest, "error", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("../images/%s-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorsMessage := gin.H{"is_uploaded": false}
		response := utils.ApiResponse("failed to upload image", http.StatusBadRequest, "error", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = ca.campaignUsecase.SaveCampaignImage(input, path)
	if err != nil {
		errorsMessage := gin.H{"is_uploaded": false}
		response := utils.ApiResponse("failed to save image", http.StatusBadRequest, "error", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errorsMessage := gin.H{"is_uploaded": true}
	response := utils.ApiResponse("success upload campaign image", http.StatusOK, "success", errorsMessage)
	c.JSON(http.StatusOK, response)
}
