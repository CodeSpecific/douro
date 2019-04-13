package controller

import "github.com/kataras/iris"

func Register(app *iris.Application) {
	registerUserController(app)
}
