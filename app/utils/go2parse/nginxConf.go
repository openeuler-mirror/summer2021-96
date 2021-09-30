package go2parse

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

func NewNginxConf(fileName string) *Config {
	cfg := config()
	cfg.loadNginxConf(fileName)
	return cfg
}
// 加载 Nginx 配置文件
func (cfg *Config) loadNginxConf(fileName string) {
	fn, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fn.Close()
	section := ""
	rd := bufio.NewReader(fn)
	for {

		data, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		line := strings.TrimSpace(string(data))

		// 若包含有 "#" 字符，分割该字符串，前部为数据，后部为注释,去掉注释部分
		if strings.Contains(line, "#"){
			line = strings.Split(line,"#")[0]
		}
		line = strings.TrimSpace(line)

		if line == "" || line[0:1] == "#" || line[0:1] == "}"{		// 若是 "#" 开头表示注释或为空行，跳过该行
			continue
		}
		if line[0] == '{' || line[len(line) - 1] == '{' { // 若是 "{" 表示一个 section
			section =line[0 : len(line) - 1]
		}
		// 将连续的空格转为一个
		line = deleteExtraSpace(line)
		cfg.parseNginxConf(line, section)
	}
}

/*
	parseNginxConf(line string,section string)
	解析文本
	line 一行文本
	section 当前 section
*/
func (cfg *Config) parseNginxConf(line string, section string) {

	if section != "" {
		if _, ok := cfg.Values[section]; !ok {
			cfg.Values[section] = make(map[string]interface{})
		}
	}
	ls := strings.SplitN(line, " ", 2) // " " 空格分割该行文字，只分割一次

	if len(ls) != 2 || strings.Contains(line,"{"){
		return
	}
	// 取出 key value
	key 	:= strings.TrimSpace(ls[0])
	value	:= strings.TrimSpace(ls[1])

	if len(value) > 0 && value[0] == '[' && value[len(value) - 1] == ']' {
		cfg.setNginxConf(section, key, strings.Split(value[1:len(value) - 1], ","))
		return
	}
	cfg.setNginxConf(section, key, value)
}

func (cfg *Config) setNginxConf(section string, key string, value interface{}) {

	if _, ok := cfg.Values[section]; !ok {
		cfg.Values[key] = value
		return
	}

	svMap := cfg.Values[section].(map[string]interface{})
	svMap[key] = value

	cfg.Values[section] = svMap
}

func deleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)      //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}