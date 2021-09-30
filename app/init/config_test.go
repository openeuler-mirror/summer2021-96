package init

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

// 测试初始化项目配置文件
func TestInitConfig(t *testing.T) {
	InitConfig()


	// 读取成功
	MySQLPath := viper.Get("default-config-file-path.MySQL")
	log.Println(MySQLPath)



}