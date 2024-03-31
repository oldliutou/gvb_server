package routers

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)
type RouterGroup struct {
    *gin.RouterGroup
}
// InitRouter初始化路由
func InitRouter() *gin.Engine{
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

 	apiRouterGroup := router.Group("/api")
	routerGroupApp := RouterGroup{RouterGroup: apiRouterGroup}
	routerGroupApp.SettingsRouter()
	
	return router

	
}


