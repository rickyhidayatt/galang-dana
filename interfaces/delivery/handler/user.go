package handler

import (
	"net/http"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/usecase"
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
		c.JSON(http.StatusBadRequest, nil)
	}
}
