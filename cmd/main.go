package main

import (
	"tgwp/initalize"
	routerg "tgwp/internal/router"
	"tgwp/log/zlog"
)

func main() {

	initalize.Init()
	// 工程进入前夕，释放资源
	defer initalize.Eve()
	routerg.RunServer()
	zlog.Infof("程序运行完成！")

}
