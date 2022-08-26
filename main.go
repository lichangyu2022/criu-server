//@title 仿真热迁移系统
// @version 1.0
// @description swagger Server api
package main

import (
	_ "criu/docs" // 这里需要引入本地已生成文档
	"criu/routes"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	log.Printf("Server started api address 0.0.0.0:swagger/index.html")
	routes.InitRouter()
}
