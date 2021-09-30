// Package init 项目初始化
package init

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"observer/app/global"
	"os"
	"path"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	projectDir, _ := os.Getwd()
	isDev := GetEnvInfo("IS_DEV")
	configFileName := path.Join(projectDir, "../../conf/application.yml")
	if isDev {
		configFileName = path.Join(projectDir, "11.config/application.dev.yml")
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 读取成功
	MySQLPath := viper.Get("default-config-file-path.MySQL")
	log.Println(MySQLPath)


	err := v.Unmarshal(&global.DefaultConfFilePath)
	if err != nil {
		fmt.Println("读取配置失败")
	}
	fmt.Println(&global.DefaultConfFilePath)
}
