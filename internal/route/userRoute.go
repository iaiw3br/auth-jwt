package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/middleware"
	"main/pkg/user"
)

func UserRoute(router *gin.Engine) *gin.Engine {
	router.Use(middleware.TokenAuthMiddleware)
	router.GET("/me", user.GetUserInformation)

	return router
}
