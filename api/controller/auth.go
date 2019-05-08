package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// authController 用户认证控制器
type authController struct{
	*iris.Application
}

// initRoute 初始化认证相关的路由
func(auth *authController)initRoute(){
	a:=auth.Party("/auth")
	a.Get("/otp/{phone:string}",hero.Handler(auth.getOtp))
}

func(auth *authController)getOtp(phone string){

}


func registerAuthController(app *iris.Application){
	 auth:=&authController{app}
	 auth.initRoute()
}

