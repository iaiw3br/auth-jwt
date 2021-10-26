package middleware

import (
	"github.com/gin-gonic/gin"
	authError "main/internal/error"
	"main/internal/token"
	"net/http"
)

func TokenAuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	isNeedUpdateTokens := token.CheckUpdateTokens(accessToken, refreshToken)

	if isNeedUpdateTokens {
		c.Set("refreshTokens", true)
	}

	username, err := token.GetUsernameFromToken(refreshToken, "SECRET_REFRESH")

	if err != nil || username == "" {
		c.JSON(http.StatusUnauthorized, authError.ErrorInvalidToken)
		c.Abort()
	}

	c.Set("username", username)

	c.Next()
}
