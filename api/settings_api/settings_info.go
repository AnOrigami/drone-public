package settings_api

import (
	"Gin_blog/global"
	"Gin_blog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//OK的测试
	//res.OK(map[string]string{
	//	"username": "ano",
	//	"sex":      "nan",
	//}, "测试成功！", c)
	//Fail的测试
	res.OKWithData(global.Config.SiteInfo, c)
}
