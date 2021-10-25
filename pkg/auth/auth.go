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
	accessToken, err := token.CreateToken(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.SetCookie("access_token", accessToken, 1500, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "success login")
}

func Logout(c *gin.Context) {
	c.SetCookie("access", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "success logout")
}
