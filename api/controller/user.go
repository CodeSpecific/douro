package controller

import (
	"github.com/CodeSpecific/douro/api/controller/internal"
	"github.com/CodeSpecific/douro/api/model"
	"github.com/CodeSpecific/douro/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// userController 用户控制器
type userController struct {
	*iris.Application
	service.UserService
	internal.Controller
}

// initRoute 初始化用户相关的路由
func (c *userController) initRoute() {
	//参数依赖注入,通过hero包实现
	c.Get("/user/{id:uint}", hero.Handler(c.getUser))
}

func (c *userController) getUser(id uint) *model.UserViewModel {
	user, err := c.GetUser(id)
	if err != nil {
		return &model.UserViewModel{Name: err.Error()}
	}
	return model.ConvertUserViewModel(user)
}

func registerUserController(app *iris.Application) {
	user := &userController{app, service.NewUserService(), &internal.BaseController{}}
	user.initRoute()
}
