package main

import (
	"github.com/CodeSpecific/douro/api/controller"
	"github.com/kataras/iris"
	"github.com/CodeSpecific/douro/infra/starter"
)

func main() {
	//使用默认中间件，自带日志和恐慌恢复组件
	app := iris.Default()
	//注册控制器
	controller.Register(app)
	//运行所有基础设施
	starter.StarterManager.Run()
	//运行api服务器
	app.Run(iris.Addr(":8080"))
}
