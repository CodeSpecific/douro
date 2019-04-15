package model

import "github.com/CodeSpecific/douro/service"

// ConvertUserViewModel 转换业务对象User为ViewModel
func ConvertUserViewModel(user *service.User) *UserViewModel {
	if user == nil {
		return nil
	}
	return &UserViewModel{
		ID:           user.ID,
		Name:         user.Name,
		Gender:       user.Gender,
		Phone:        user.Phone,
		Avatar:       user.Avatar,
		RegisterMode: user.RegisterMode,
		CreateTime:   user.CreateTime,
	}
}
