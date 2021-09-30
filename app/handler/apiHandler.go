// Package handler 数据请求相关 API Handler
package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"observer/app/utils"
	"observer/app/utils/go2parse"
)

// ConfFilePath 配置文件路径 //TODO 后期通过专门的配置文件来指定
var (
	ConfFilePath = map[string]string{
		"MySQL":"/etc/mysql/mysql.conf.d/mysqld.cnf",
		"Redis":"/etc/redis/redis.conf",
		"Crontab": "/etc/crontab",
		"Nginx": "/etc/nginx/nginx.conf",
		"iptables": "/etc/sysconfig/iptables",
	}
)

// ConfigurationReadByObserver 读取解析后的配置数据
func ConfigurationReadByObserver(context *gin.Context) {

	// 获取当前配置状态（配置目标）
	currentState := context.Query("current_state")
	log.Printf("current_state: %s\n", currentState)

	path := context.DefaultQuery("configuration_file_path",ConfFilePath[currentState])

	i := go2parse.New(path)
	//log.Println(i)

	// 页面及初始变量
	context.JSON(http.StatusOK,gin.H{
		"current_state":currentState,
		"conf":i,
		"success": true,
	})
}

// ConfigurationWriteByObserver 写入更新的配置数据
func ConfigurationWriteByObserver(context *gin.Context) {

	// 通过使用 ioutil 来读取请求体
	//data, _ := ioutil.ReadAll(context.Request.Body)
	//log.Printf("ctx.Request.body: %v\n", string(data))
	//var dat map[string]interface{}
	//if err := json.Unmarshal(data, &dat); err == nil {
	//	log.Println(dat)
	//	log.Println(dat["port"])
	//} else {
	//	log.Fatalln(err)
	//}

	json := make(map[string]interface{})

	context.BindJSON(&json)
	//log.Printf("BINDJSON: %v",&json)

	// 获取当前操作状态
	currentState := json["current_state"]
	log.Printf("CURRENT_STATE %s",currentState)

	// 通过当前状态来调用不同的方法来写入配置
	switch currentState {
	case "MySQL":
		mySqlWrite(json)
	case "Redis":
		redisWrite(json)
	default:
		log.Println("未匹配！")
	}

	//返回结果
	context.JSON(http.StatusOK, gin.H{
		"success":         true,
	})
}
func mySqlWrite(json map[string]interface{}) {
	//和原来配置 map 对比，将有差异配置项的写入文件
	fileName := ConfFilePath["MySQL"]
	i := go2parse.New(fileName)

	for _,v := range i.Values{
		for ik,iv := range v.(map[string]interface{}){
			jIk := json[ik]
			if jIk != iv {
				log.Printf("配置更新 ( %s ): %s => %s \n",ik,iv,json[ik])
				utils.ReplaceText(fileName,ik,json[ik].(string)," = ")
			}
		}
	}
}
func redisWrite(json map[string]interface{}) {
	//和原来配置 map 对比，将有差异配置项的写入文件
	fileName := ConfFilePath["Redis"]
	i := go2parse.New(fileName)

	for k,v := range i.Values{
			jIk := json[k]
			if jIk != v {
				log.Printf("配置更新 ( %s ): %s => %s \n",k,v,json[k])
				utils.ReplaceText(fileName,k,json[k].(string)," ")
			}
	}
}

// ConfigurationFileRead  配置文件读取
func ConfigurationFileRead(context *gin.Context)  {
	isRead := true //读取文件是否成功
	// 获取当前配置状态（配置目标）
	currentState := context.Query("current_state")
	log.Printf("current_state: %s\n", currentState)

	path := context.DefaultQuery("configuration_file_path",ConfFilePath[currentState])
	//文件读取任务是将文件内容读取到内存中。
	info, err := ioutil.ReadFile(path)
	if err!=nil{
		log.Println(err)
		isRead =false
	}
	//log.Println(info)
	result:=string(info)

	//返回结果
	context.JSON(http.StatusOK, gin.H{
		"file_name": ConfFilePath[currentState], // TODO: 通过正则来截取 path 中的文件名
		"content": result,
		"success": isRead,
	})
}

// ConfigurationFileWrite 写配置文件
// @Param context *gin.Context
// @return nil
func ConfigurationFileWrite(context *gin.Context){

	// 获取当前配置状态（配置目标）
	currentState := context.PostForm("current_state")
	log.Println(currentState)

	isWrite :=true // Flag 写入文件是否成功
	//需要写入到文件的内容
	path := context.DefaultQuery("configuration_file_path",ConfFilePath[currentState])
	updatedContent := context.PostForm("updated_content")

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

