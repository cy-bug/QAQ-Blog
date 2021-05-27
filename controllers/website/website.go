package website

import (
	"QAQ-Blog/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadIndex(context *gin.Context) {
	page := context.Param("page")
	context.HTML(http.StatusOK, "index.html", gin.H{
		"name": config.Config.GetString("website.username"),
		"page": page,
	})
}

func LoadAbout(context *gin.Context) {
	context.HTML(http.StatusOK, "about.html", gin.H{
		"name":      config.Config.GetString("website.username"),
		"aboutinfo": config.Config.GetString("website.aboutinfo"),
		"mail":      config.Config.GetString("website.mail"),
		"qq":        config.Config.GetString("website.qq"),
	})
}
