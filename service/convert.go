package service

import (
	"github.com/CodeSpecific/douro/core/entity"
)

// convertToUser 根据User_Profile和User_Auth转化为业务对象User
func convertToUser(profile *entity.User_Profile, auth *entity.User_Auth) *User {
	if profile == nil && auth == nil {
		return nil
	}
	user := &User{}
	if profile != nil {

		user.ID = profile.ID
		user.Name = profile.Name
		user.Phone = profile.Phone
		user.Avatar = profile.Avatar
		user.ThirdPartyID = profile.Third_Party_ID
		user.CreateTime = profile.Create_Time

		//类型断言并进行强制转换
		if gender, ok := interface{}(profile.Gender).(Gender); ok {
			user.Gender = gender
		}

		if mode, ok := interface{}(profile.Register_Mode).(RegisterMode); ok {
			user.RegisterMode = mode
		}
	}

	if auth != nil {
		user.EncryptPwd = auth.Encrypt_Pwd
	}

	return user
}
