package routers

import (
	"Gin_blog/api"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings", settingsApi.SettingsInfoView)
	
}
