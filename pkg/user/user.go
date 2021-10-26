package user

import (
	"github.com/gin-gonic/gin"
	authError "main/internal/error"
	"main/internal/token"
	"net/http"
)

func GetUserInformation(c *gin.Context) {
	username, _ := c.Get("username")
	if usernameString, ok := username.(string); ok {
		_, ok := c.Get("refreshTokens")
		if ok {
			tokens, err := token.CreateTokens(usernameString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, authError.ErrorInvalidToken)
			}
			c.SetCookie("access_token", tokens["accessToken"], 60, "/", "localhost", false, true)
			c.SetCookie("refresh_token", tokens["refreshToken"], 3600, "/", "localhost", false, true)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": username,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, authError.ErrIncomingData)
}
