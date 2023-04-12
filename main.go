package main

import (
	"Gin_blog/core"
	"Gin_blog/flag"
	"Gin_blog/global"
	"Gin_blog/routers"
)

func main() {
	//读取settings.yaml配置文件
	core.InitConf()

	//初始化日志
	global.Log = core.InitLogger()

	//链接数据库
	global.DB = core.InitGorm()

	//路由测试
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("运行在%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
