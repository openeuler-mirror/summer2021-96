package utils

import (
	"fmt"
	"testing"
)
// 获取当前项目所在路径
func TestCurrentPath(t *testing.T)  {
	fmt.Println("getTmpDir（当前系统临时目录） = ", getTmpDir())
	fmt.Println("getCurrentAbPathByExecutable（仅支持go build） = ", getCurrentAbPathByExecutable())
	fmt.Println("getCurrentAbPathByCaller（仅支持go run） = ", getCurrentAbPathByCaller())
	fmt.Println("getCurrentAbPath（最终方案-全兼容） = ", GetCurrentAbPath())
}
