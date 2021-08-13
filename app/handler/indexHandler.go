package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"observer/app/utils"
	"observer/app/utils/go2parse"
)

/*路径请求相关 Handler*/

// Index 项目起始页
func Index(context *gin.Context) {
	context.HTML(http.StatusOK,"index.gohtml",gin.H{
		"title": "Observer，帮您更优雅的配置~",
	})
}

// Main 后台主页
func Main(context *gin.Context){
	context.HTML(http.StatusOK,"main.gohtml",gin.H{
		"welcome_info": "Observer，帮您更优雅的配置~",
	})
}

// ConfigMysql Mysql 配置页
func ConfigMysql(context *gin.Context) {
	// 页面及初始变量
	context.HTML(http.StatusOK,"c_mysql.gohtml",gin.H{
		"current_state":"MySQL",
	})
}
// ConfigRedis Redis 配置页
func ConfigRedis(context *gin.Context) {
	fileName := "H:\\redis.conf"
	i := go2parse.New(fileName)
	log.Println(i)

	fileContent := utils.ReadFile(fileName)
	log.Println(fileContent)

	context.HTML(http.StatusOK,"c_redis.gohtml",gin.H{
		"redis_config":i,
		"file_content":fileContent,
	})
}

// ConfigCrontab  Crontab 配置页
func ConfigCrontab(context *gin.Context) {
	//fileName := "H:\\crontab"
	//i := go2parse.New(fileName)
	//log.Println(i)
	//
	//fileContent := utils.ReadFile(fileName)
	//log.Println(fileContent)

	context.HTML(http.StatusOK,"c_crontab.gohtml",gin.H{
		//"crontab_config":i,
		//"file_content":fileContent,
	})
}