package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"main/internal/service"
	"main/internal/token"
	"net/http"
)

func Login(c *gin.Context) {

	err := service.CheckValidateUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	username := viper.GetString("USERNAME")
	tokens, err := token.CreateTokens(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.SetCookie("access_token", tokens["accessToken"], 60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens["refreshToken"], 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, "success login")
}

func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "success logout")
}
