package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"observer/app/utils/go2parse"
)

/* API 请求相关 Handler */

// ConfigurationReadByObserver 读取解析后的配置数据
func ConfigurationReadByObserver(context *gin.Context) {

	path := context.DefaultQuery("configuration_file_path","H://mysql.cnf")

	i := go2parse.New(path)
	log.Println(i)

	// 页面及初始变量
	context.JSON(http.StatusOK,gin.H{
		"current_state":"MySQL",
		"config":i,
		"success": true,
	})
}

// ConfigurationFileRead  配置文件读取 */
func ConfigurationFileRead(context *gin.Context)  {

	isRead := true //读取文件是否成功
	// 获取当前配置状态（配置目标）
	currentState := context.Query("current_state")
	log.Printf("current_state: %s\n", currentState)

	path := context.DefaultQuery("configuration_file_path","H://mysql.cnf")
	//文件读取任务是将文件内容读取到内存中。
	info, err := ioutil.ReadFile(path)
	if err!=nil{
		log.Println(err)
		isRead =false
	}
	log.Println(info)
	result:=string(info)

	//返回结果
	context.JSON(http.StatusOK, gin.H{
		"file_name": "mysql.cnf", // TODO: 通过正则来截取 path 中的文件名
		"content": result,
		"success": isRead,
	})
}

// ConfigurationFileWrite 将内容写入文件
func ConfigurationFileWrite(context *gin.Context){
	isWrite :=true //写入文件是否成功
	//需要写入到文件的内容
	path := context.DefaultQuery("configuration_file_path","H://mysql.cnf")
	updatedContent := context.PostForm("updated_content")
	currentState := context.PostForm("current_state")
	log.Println(currentState)

	d1 := []byte(updatedContent)
	err := ioutil.WriteFile(path, d1, 0644)

	if err!=nil{
		isWrite =false
	}
	//返回结果
	context.JSON(http.StatusOK, gin.H{
		"success":         isWrite,
		"updatedContent": updatedContent,
	})
}

