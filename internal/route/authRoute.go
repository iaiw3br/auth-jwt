package route

import (
	"github.com/gin-gonic/gin"
	"main/pkg/auth"
)

func AuthRoute(router *gin.Engine) *gin.Engine {
	router.POST("/login", auth.Login)
	router.POST("/logout", auth.Logout)

	return router
}
