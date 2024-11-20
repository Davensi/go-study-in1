package main

import (
	"common/config"
	"common/metrics"
	"user/app"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

var conf = flag.String("conf", "./application.yml", "conf file")

func main() {

	// 1.加载配置
	flag.Parse()
	config.IninConfig(*conf)

	// 2.启动监控
	go func() {
		if err := metrics.Serve(fmt.Sprintf("0.0.0.0:%d", config.Conf.MetricPort)); err != nil {
			print("实时监控错误--->", err)
		}
	}()
	// 3.启动grpc服务
	if err := app.Run(context.Background()); err != nil {
		log.Printf("grpc服务错误--->", err)
		os.Exit(-1)
	} 
}
