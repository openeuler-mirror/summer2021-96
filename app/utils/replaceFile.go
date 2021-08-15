package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// ReplaceText 替换指定配置文件中的文本
// @param path 要操作文件的路径
// @param key  配置文件中要替换的配置项名称
// @param value 配置文件中要替换的配置项的值
// @param delimiter 分隔符（不同的配置文件所用的分隔符以相同） //TODO: 使用该方法的不便之处，需要考虑不同配置文件配置项的分隔符
// @return nil
// TODO: BUG1:替换时该行的长度 pos 为原始配置项对应行的长度，导致在更新配置时若长度不相同，则会对配置文件修改错误
func ReplaceText(path string,key string, value string, delimiter string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Println("open file filed.", err)
		return
	}
	//defer关闭文件
	defer file.Close()
	//获取文件大小
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	log.Println("file size:", size) //读取文件内容到io中
	reader := bufio.NewReader(file)
	pos := int64(0) // TODO: BUG1


	for { //读取每一行内容
		line, err := reader.ReadString('\n')
		if err != nil { //读到末尾
			if err == io.EOF {
				log.Println("File read ok!")
				break
			} else {
				log.Println("Read file error!", err)
				return
			}
		}
		//log.Println(line) //根据关键词覆盖当前行
		if strings.Contains(line, key) {
			bytes := []byte(key + delimiter + value + "\n")
			file.WriteAt(bytes, pos)
		}
		//每一行读取完后记录位置
		pos += int64(len(line))
	}
}