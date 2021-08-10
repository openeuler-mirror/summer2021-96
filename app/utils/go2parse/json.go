package go2parse

import (
	"encoding/json"
	"io/ioutil"
)

func NewJson(fileName string, v interface{}) *Config {
	cfg := config()

	if v == nil {
		cfg.LoadJson(fileName, &cfg.Values)
	} else {
		cfg.LoadJson(fileName, &v)
	}

	return cfg
}

func (cfg *Config) LoadJson (fileName string, v interface{}) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		panic(err)
	}

	if err := json.Unmarshal([]byte(data), v); err != nil {
		panic(err)
	}
}