package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	initAPIRouters(r)
	initWebsiteRouters(r)
}
