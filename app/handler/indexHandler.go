// Package handler 页面请求相关 API Handler
package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 项目起始页 TODO: 已经弃用，后期用作登录页面
func Index(context *gin.Context) {
	context.HTML(http.StatusOK,"main.gohtml",gin.H{
		"title": "Observer，帮您更优雅的配置~",
	})
}

// Main 后台主页
func Main(context *gin.Context){
	context.HTML(http.StatusOK,"main.gohtml",gin.H{
		"welcome_info": "Observer，帮您更优雅的配置~",
	})
}

// ConfigurationPage 标准配置页面 (通过参数 "current_state",来判断当前配置状态)
func ConfigurationPage(context *gin.Context) {
	currentState := context.Query("cs")
	context.HTML(http.StatusOK,"configuration.gohtml",gin.H{
		"current_state": currentState,
	})
}