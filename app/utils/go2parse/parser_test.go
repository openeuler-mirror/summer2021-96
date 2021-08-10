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

	i := New("../../../statics/config_file/mysql/mysql_simple.cnf")
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

	i := New("../../../statics/config_file/redis/redis.conf")
	fmt.Println(i)

	for k, v := range i.Values {
		fmt.Printf("key: %s <==> value: %s \n",k,v)
	}

	fmt.Println(i.Values["lua-time-limit"])
	fmt.Println(i.Get("appendfilename"))
}