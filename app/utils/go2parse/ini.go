package go2parse

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func NewIni(fileName string) *Config {
	cfg := config()
	cfg.loadIni(fileName)

	return cfg
}

func (cfg *Config) loadIni(fileName string) {

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
		if line == "" || line[0:1] == "#" {
			continue
		}

		if line[0] == '[' && line[len(line) - 1] == ']' {
			section = line[1 : len(line) - 1]
		}

		cfg.parse(line, section)
	}
}

func (cfg *Config) parse(line string, section string) {

	if section != "" {
		if _, ok := cfg.Values[section]; !ok {
			cfg.Values[section] = make(map[string]interface{})
		}
	}

	ls := strings.Split(line, "=")
	if len(ls) != 2 {
		return
	}

	key 	:= strings.TrimSpace(ls[0])
	value	:= strings.TrimSpace(ls[1])

	if len(value) > 0 && value[0] == '[' && value[len(value) - 1] == ']' {
		cfg.set(section, key, strings.Split(value[1:len(value) - 1], ","))
		return
	}

	cfg.set(section, key, value)
}

func (cfg *Config) set(section string, key string, value interface{}) {

	if _, ok := cfg.Values[section]; !ok {
		cfg.Values[key] = value
		return
	}

	svMap := cfg.Values[section].(map[string]interface{})
	svMap[key] = value

	cfg.Values[section] = svMap
}