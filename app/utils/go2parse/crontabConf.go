package go2parse

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func NewCrontabConf(fileName string) *Config {
	cfg := config()
	cfg.loadCrontabConf(fileName)
	return cfg
}
// 加载 Crontab 配置文件
func (cfg *Config) loadCrontabConf(fileName string) {
	fn, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fn.Close()
	section := ""
	rd := bufio.NewReader(fn)
	var taskNum int
	for {

		data, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		line := strings.TrimSpace(string(data)) // 去除首尾


		if line == "" || line[0:1] == "#"{		// 若是 "#" 开头表示注释或为空行，跳过该行
			continue
		}
		if line[0] == '{' || line[len(line) - 1] == '{' { // 若是 "{" 表示一个 section
			section =line[0 : len(line) - 1]
		}
		// 如果首字符不是字母
		if !unicode.IsLetter(rune(line[0])){
			section = "task" + strconv.Itoa(taskNum)
			taskNum ++

		}
		// 将连续的空格转为一个
		line = deleteExtraSpace(line)
		cfg.parseCrontabConf(line, section)
	}
}

/*
	parseCrontabConf(line string,section string)
	解析文本
	line 一行文本
	section 当前 section
*/
func (cfg *Config) parseCrontabConf(line string, section string) {

	if section != "" {
		if _, ok := cfg.Values[section]; !ok {
			cfg.Values[section] = make(map[string]interface{})
		}
	}
	var ls []string
	var key string
	var value string

	if unicode.IsLetter(rune(line[0])) { // 如果是字母，则说明是 Key-Value配置
		ls = strings.SplitN(line,"=",2)
		// 取出 key value
		key 	= strings.TrimSpace(ls[0])
		value	= strings.TrimSpace(ls[1])

		cfg.setCrontabConf(section,key,value)
	}else {
		//ls = strings.SplitN(line," ",7)
		// 取出 key value
		//cron 	= strings.TrimSpace(ls[0] + " " + ls[1] + " " + ls[2] + " " + ls[3] + " " + ls[4])
		//user	= strings.TrimSpace(ls[5])
		//command	= strings.TrimSpace(ls[6])

		cfg.Values[section] = make(map[string]interface{})
		cfg.setCrontabConf(section,section,line)

	}

	//fmt.Println(cron)
	//fmt.Println(user)
	//fmt.Println(command)

	//if len(ls) != 2 || strings.Contains(line,"{"){
	//	return
	//}



}

func (cfg *Config) setCrontabConf(section string, key string, value interface{}) {

	if _, ok := cfg.Values[section]; !ok {
		cfg.Values[key] = value
		return
	}



	ls := strings.SplitN(fmt.Sprintf("%v",value)," ",7)
	cron 	:= strings.TrimSpace(ls[0] + " " + ls[1] + " " + ls[2] + " " + ls[3] + " " + ls[4])
	user	:= strings.TrimSpace(ls[5])
	command	:= strings.TrimSpace(ls[6])

	svMap := cfg.Values[section].(map[string]interface{})
	svMap["cron"] = cron
	svMap["user"] = user
	svMap["command"] = command



	cfg.Values[section] = svMap
}
