package go2parse

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func NewCnf(fileName string) *Config {
	cfg := config()
	cfg.loadCnf(fileName)

	return cfg
}

func (cfg *Config) loadCnf(fileName string) {

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
		if line == "" || line[0:1] == "#" {		// 若是 "#" 开头表示注释或为空行，跳过该行
			continue
		}

		if line[0] == '[' && line[len(line) - 1] == ']' { // 若是 "[]" 表示一个 section
			section = line[1 : len(line) - 1]
		}

		cfg.parseCnf(line, section)
	}
}

/*
	parseCnf(line string,section string)
	解析文本
	line 一行文本
	section 当前 section
*/
func (cfg *Config) parseCnf(line string, section string) {

	if section != "" {
		if _, ok := cfg.Values[section]; !ok {
			cfg.Values[section] = make(map[string]interface{})
		}
	}

	ls := strings.Split(line, "=") // "=" 分割该行文字
	if len(ls) != 2 {
		return
	}

	// 取出 key value
	key 	:= strings.TrimSpace(ls[0])
	value	:= strings.TrimSpace(ls[1])

	if len(value) > 0 && value[0] == '[' && value[len(value) - 1] == ']' {
		cfg.setCnf(section, key, strings.Split(value[1:len(value) - 1], ","))
		return
	}

	cfg.setCnf(section, key, value)
}

func (cfg *Config) setCnf(section string, key string, value interface{}) {

	if _, ok := cfg.Values[section]; !ok {
		cfg.Values[key] = value
		return
	}

	svMap := cfg.Values[section].(map[string]interface{})
	svMap[key] = value

	cfg.Values[section] = svMap
}