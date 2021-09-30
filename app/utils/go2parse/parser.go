package go2parse

import (
	"log"
	"path"
	"reflect"
)

type Config struct {
	Values  map[string]interface{}
}

func New(fileName string) * Config {

	// 通过 完整文件名来判断（文件名+扩展名）
	fullName := path.Base(fileName)
	log.Println(fullName)

	// 通过扩展名来判断
	//switch path.Ext(fileName) {

	switch path.Base(fileName) {
	case "redis.conf":
		return NewRedisConf(fileName)
	case "mysqld.cnf":
		return NewCnf(fileName)
	case "nginx.conf":
		return NewNginxConf(fileName)
	case "nginx_simple.conf":
		return NewNginxConf(fileName)
	case "crontab":
		return NewCrontabConf(fileName)
	case ".ini":
		return NewIni(fileName)
	case "conf.json":
		return NewJson(fileName, nil)
	case ".xml":
		return NewXml(fileName, nil)
	case ".yaml":
		return NewYaml(fileName)
	default:
		panic("conf file error")
	}

	return nil
}

func config() *Config {
	return &Config{Values: make(map[string]interface{})}
}

func (cfg *Config) toSlice(vs interface{}) []string {

	v := reflect.ValueOf(vs)
	if v.Kind() != reflect.Slice {
		panic("to slice error")
	}

	ret := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = v.Index(i).Interface().(string)
	}

	return ret
}

func (cfg *Config) Get (keys ...string) (v interface{}){

	target := cfg.Values

	for _, key := range keys {
		if v, ok := target[key]; ok {

			switch reflect.TypeOf(v).Kind() {
				case reflect.Map:
					target = v.(map[string]interface{})
					break
				case reflect.Slice:
					return cfg.toSlice(v)
				default:
					return v
			}

		} else {
			return nil
		}
	}

	return target
}