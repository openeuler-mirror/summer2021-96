package go2parse

import (
	"fmt"
	"log"
	"testing"
)
// mysql.cnf 文件解析
func TestMysqlCnf(t *testing.T) {
	log.Println("-----------------------------------------")
	log.Println("mysql.cnf")

	//i := New("../../../statics/config_file/mysql/mysql_simple.cnf")
	i := New("/etc/mysql/mysql.conf.d/mysqld.cnf")
	//i := New("H:\\ProjectWarehouse\\GoLandProjects\\observer\\statics\\config_file\\mysql\\mysql_simple.cnf")
	fmt.Printf("i : %s\n",i)
	fmt.Printf("i.Values : %s\n",i.Values)
	fmt.Println((i.Get("mysqld")).(map[string]interface{}))

	for k,v := range i.Values{
		fmt.Printf("section: k = %s v= %s\n",k,v)
		for ik,iv := range v.(map[string]interface{}){
			fmt.Printf("item: ik = %s iv = %s\n",ik,iv)
		}
	}

	// 遍历单独 section
	//for k,v := range (i.Get("mysqld")).(map[string]interface{}){
	//	fmt.Printf("key: %s <==> value: %s\n",k,v)
	//}

}
// redis.conf 文件解析
func TestRedisConf(t *testing.T) {
	log.Println("-----------------------------------------")
	log.Println("redis.conf")

	//i := New("../../../statics/config_file/redis/redis.conf")
	i := New("../../../statics/config_file/redis/redis.conf")
	fmt.Println(i)

	for k, v := range i.Values {
		fmt.Printf("key: %s <==> value: %s \n",k,v)
	}

	fmt.Println(i.Values["lua-time-limit"])
	fmt.Println(i.Get("appendfilename"))
}
// nginx.conf 文件解析
func TestNginxConf(t *testing.T) {
	log.Println("-----------------------------------------")
	log.Println("nginx.conf")

	//i := New("../../../statics/config_file/redis/redis.conf")
	//i := New("../../../statics/config_file/nginx/nginx.conf")
	i := New("../../../statics/config_file/nginx/nginx_simple.conf")
	fmt.Println(i)

	//for k, v := range i.Values {
	//	fmt.Printf("%s : %s \n",k,v)
	//}
	for k, v := range i.Values {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")

		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}
// crontab 文件解析
func TestCrontabConf(t *testing.T) {
	log.Println("-----------------------------------------")
	log.Println("crontab.conf")

	//i := New("../../../statics/config_file/redis/redis.conf")
	//i := New("../../../statics/config_file/nginx/nginx.conf")
	i := New("../../../statics/config_file/crontab/crontab")
	fmt.Println(i)

	for k, v := range i.Values {
		fmt.Printf("%s : %s \n",k,v)
	}
	//for k, v := range i.Values {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Println(k, "is string", vv)
	//	case float64:
	//		fmt.Println(k, "is float", int64(vv))
	//	case int:
	//		fmt.Println(k, "is int", vv)
	//	case []interface{}:
	//		fmt.Println(k, "is an array:")
	//		for i, u := range vv {
	//			fmt.Println(i, u)
	//		}
	//	case nil:
	//		fmt.Println(k, "is nil", "null")
	//	case map[string]interface{}:
	//		fmt.Println(k, "is an map:")
	//	default:
	//		fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
	//	}
	//}
}