package go2parse

import (
	"path"
	"reflect"
)

type Config struct {
	Values  map[string]interface{}
}

func New(fileName string) * Config {

	switch path.Ext(fileName) {
	case ".conf":
		return NewRedisConf(fileName)
	case ".cnf":
		return NewCnf(fileName)
	case ".ini":
		return NewIni(fileName)
	case ".json":
		return NewJson(fileName, nil)
	case ".xml":
		return NewXml(fileName, nil)
	case ".yaml":
		return NewYaml(fileName)
	default:
		panic("config file error")
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