package main

import (
	"observer/app/route"
)

// 程序入口
func main() {
	router := route.SetupRouter()
	_ = router.Run()
}
