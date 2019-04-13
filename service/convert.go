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

		//强制类型转换
		user.Gender = Gender(profile.Gender)
		user.RegisterMode = RegisterMode(profile.Register_Mode)
	}

	if auth != nil {
		user.EncryptPwd = auth.Encrypt_Pwd
	}

	return user
}
