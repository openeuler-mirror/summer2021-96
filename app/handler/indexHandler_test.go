package handler

import (
	"log"
	"observer/app/utils/go2parse"
	"testing"
)

// 测试 mysql.cnf

func TestConfigMysql(t *testing.T) {
	fileName := "H:\\mysql.cnf"
	i := go2parse.New(fileName)
	log.Println(i)
}