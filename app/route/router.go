package route

import (
	"github.com/gin-gonic/gin"
	"observer/app/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	/*通过不同的 mode 加载 templates*/
	if mode := gin.Mode(); mode == gin.TestMode{
		router.LoadHTMLGlob("./../templates/*")
	}else {
		router.LoadHTMLGlob("templates/*")
		//router.LoadHTMLGlob("templates/**/*")
	}

	/*静态资源路由*/
	router.Static("/statics","./statics")
	// favicon.ico
	router.StaticFile("/favicon.ico","./favicon.ico")

	/*路由分组*/
	index := router.Group("/")
	{
		index.GET("",handler.Main)
		index.GET("main",handler.Main)

		configGroup := router.Group("config")
		{
			configGroup.GET("/mysql",handler.ConfigMysql)
			configGroup.GET("/redis",handler.ConfigRedis)
			configGroup.GET("/crontab",handler.ConfigCrontab)
		}
	}
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/configuration_file",handler.ConfigurationFileRead)
		apiGroup.POST("/configuration_file",handler.ConfigurationFileWrite)
		apiGroup.GET("/configuration_file/observer",handler.ConfigurationReadByObserver)
	}


	return router
}