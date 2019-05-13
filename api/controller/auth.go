package controller

import (
	"github.com/CodeSpecific/douro/api/model"
	"github.com/CodeSpecific/douro/kit"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// authController 用户认证控制器
type authController struct {
	*iris.Application
}

// initRoute 初始化认证相关的路由
func (auth *authController) initRoute() {
	a := auth.Party("/auth")
	a.Get("/otp/{phone:string}", hero.Handler(auth.getOtp))
}

func (auth *authController) getOtp(phone string) interface{} {
	if !kit.Validate.IsPhone(phone) {
		return model.ResponseCommonModel.Error("输入的手机号码不正确，请检查")
	}
	// 1.这里需要对每次的调用进行时间限制，例如一分钟
	// 2.使用redis存储电话号码和随机数，并设置过期时间为1分钟
	// opt 匿名结构体对象
	opt := struct {
		Phone string
		Code  string
	}{
		phone,
		"123456",
	}
	return model.ResponseCommonModel.Success(opt)
}

func registerAuthController(app *iris.Application) {
	auth := &authController{app}
	auth.initRoute()
}
