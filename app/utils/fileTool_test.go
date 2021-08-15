package utils

import (
	"fmt"
	"testing"
)

// TestFileRead 测试文件读取
func TestFileRead(t *testing.T) {
	fmt.Println(ReadFile("H:\\redis.conf"))
}