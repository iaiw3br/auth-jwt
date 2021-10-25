package middleware

import (
	"github.com/gin-gonic/gin"
	"main/internal/auth"
	"net/http"
)

func TokenAuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	_, err = auth.ValidateToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	c.Next()
}
