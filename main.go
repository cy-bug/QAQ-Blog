package main

import (
	"QAQ-Blog/config"
	_ "QAQ-Blog/config"
	"QAQ-Blog/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routers.InitRouters(engine)
	_ = engine.Run(":" + config.Config.GetString("website.port"))
}
