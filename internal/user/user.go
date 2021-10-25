package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInformation(c *gin.Context) {
	// TODO: Добавить информацию о пользователе
	c.JSON(http.StatusOK, "There are must be full personal information.")
}
