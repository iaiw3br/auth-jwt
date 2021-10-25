package route

import "github.com/gin-gonic/gin"

func New(router *gin.Engine) {
	AuthRoute(router)
	UserRoute(router)
}
