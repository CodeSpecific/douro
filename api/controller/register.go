package controller

import "github.com/kataras/iris"

// Register 注册控制器
func Register(app *iris.Application) {
	registerUserController(app)
}
