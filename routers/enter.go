package routers

import (
	"Gin_blog/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//路由分组
	apiGroup := router.Group("api")
	//系统配置api
	SettingsRouter(apiGroup)

	return router
}
