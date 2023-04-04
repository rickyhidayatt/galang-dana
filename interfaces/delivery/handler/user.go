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
