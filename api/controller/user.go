package controller

import (
	"github.com/CodeSpecific/douro/api/controller/internal"
	"github.com/CodeSpecific/douro/api/model"
	"github.com/CodeSpecific/douro/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

type userController struct {
	*iris.Application
	service.UserService
	internal.Controller
}

func (c *userController) initRoute() {
	//参数依赖注入,通过hero包实现
	c.Get("/user/{id:uint}", hero.Handler(c.getUser))
}

func (c *userController) getUser(id uint) *model.UserViewModel {
	user, _ := c.GetUser(id)
	return model.ConvertUserViewModel(user)
}

func registerUserController(app *iris.Application) {
	user := &userController{app, service.NewUserService(), &internal.BaseController{}}
	user.initRoute()
}
