package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/interfaces/delivery/auth"
	"github.com/galang-dana/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Service, userService usecase.UserUseCase) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenSTR := strings.Split(authHeader, " ")
		if len(tokenSTR) == 2 {
			tokenString = tokenSTR[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := utils.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := utils.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := claim["user_id"].(string)
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := utils.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
