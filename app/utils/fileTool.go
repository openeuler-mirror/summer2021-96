package utils

import (
	"fmt"
	"io"
	"os"
)

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ReadFile(fileName string) string{
	// 打开文件
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败。", err)
		return "err"
	}
	defer fp.Close()

	buf := make([]byte, 1024)
	// 读取文件 (块读取)
	for {
		// 循环读取文件
		n, err2 := fp.Read(buf)
		if err2 == io.EOF {  // io.EOF表示文件末尾
			fmt.Println("文件读取结束")
			break
		}
		return string(buf[:n])
	}
	return ""
}