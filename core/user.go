package core

import "github.com/CodeSpecific/douro/core/entity"

// UserCore DAO对象
type UserCore interface {
	GetUserProfile(id uint) *entity.User_Profile
	GetUserAuth(id uint) *entity.User_Auth
}

type userCore struct {
}

func (c *userCore) GetUserProfile(id uint) *entity.User_Profile {
	return &entity.User_Profile{
		ID:     id,
		Name:   "Vitamin",
		Gender: 1,
	}
}

func (c *userCore) GetUserAuth(id uint) *entity.User_Auth {
	return &entity.User_Auth{}
}

func NewUserCore() UserCore {
	return &userCore{}
}
