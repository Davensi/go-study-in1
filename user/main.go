package main

import (
	"flag"
	"fmt"
	"common/config"
)

var conf = flag.String("conf","./application.yml","conf file")
func main() {
	// 1.加载配置
	flag.Parse()
	config.IninConfig(*conf)
	fmt.Println("加载配置成功",config.Conf)
	// 2.启动监控
	// 3.启动grpc服务

	select {
		
	}
}