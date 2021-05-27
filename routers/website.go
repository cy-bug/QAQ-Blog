package routers

import (
	"QAQ-Blog/controllers/website"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initWebsiteRouters(engine *gin.Engine) {
	engine.LoadHTMLGlob("./static/site/*")
	engine.StaticFS("/static", http.Dir("./static"))

	// 主页重定向
	engine.GET("/", func(context *gin.Context) {
		context.Request.URL.Path = "/website/index/1"
		engine.HandleContext(context)
	})

	// 添加网页路由
	websiteRouters := engine.Group("/website")
	websiteRouters.GET("/index/:page", website.LoadIndex)
	websiteRouters.GET("/about", website.LoadAbout)
}
