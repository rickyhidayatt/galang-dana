package handler

import (
	"net/http"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/utils"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase usecase.UserUseCase
}

func NewUserHandler(userUC usecase.UserUseCase) *userHandler {
	return &userHandler{userUC}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input input.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := utils.ApiResponse("Server Error", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	registerUser, err := h.userUsecase.Register(input)
	if err != nil {
		response := utils.ApiResponse("failed register your account", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.ApiResponse("successfully register your account", http.StatusOK, "success", registerUser)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input input.LoginUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidatorError(err)
		errorsMessage := gin.H{"error": errors}

		response := utils.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userUsecase.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := utils.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := utils.ApiResponse("successfully login", http.StatusOK, "success", loginUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input input.CheckEmail

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidatorError(err)
		response := utils.ApiResponse("email checking failed", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	emailOK, err := h.userUsecase.EmailAvaliable(input)
	if err != nil {
		errors := gin.H{"error": "server error, u can use your email"}
		response := utils.ApiResponse("email checking failed", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_avaliable": emailOK,
	}

	metaMessage := "Email has been registered"

	if emailOK {
		metaMessage = "Email is Avaliable"
	}

	response := utils.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := utils.ApiResponse("failed upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	folderSave := "../images/" + file.Filename

	err = c.SaveUploadedFile(file, folderSave)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := utils.ApiResponse("failed upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId := "c7a626ab48cc419399b5c662e6a9043"

	_, err = h.userUsecase.SaveAvatar(userId, folderSave)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := utils.ApiResponse("failed uploaded avatar, id not registered", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := utils.ApiResponse("success uploaded avatar images", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
