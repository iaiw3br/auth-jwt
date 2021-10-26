package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	authError "main/internal/error"
	"main/internal/model"
)

func CheckValidateUser(c *gin.Context) error {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		return errors.New(authError.ErrInvalidJson)
	}

	username := viper.GetString("USERNAME")
	password := viper.GetString("PASSWORD")

	if username != user.Username || password != user.Password {
		return errors.New(authError.ErrInvalidLoginData)
	}

	return nil
}
