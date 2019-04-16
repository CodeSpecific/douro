package core

import (
	"github.com/CodeSpecific/douro/core/entity"
)

// UserCore DAO对象
type UserCore interface {
	GetUserProfile(id uint) (*entity.User_Profile, error)
	GetUserAuth(id uint) (*entity.User_Auth, error)
}

type userCore struct {
}

func (c *userCore) GetUserProfile(id uint) (*entity.User_Profile, error) {
	user := entity.User_Profile{}
	has, err := db.Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &user, nil
}

func (c *userCore) GetUserAuth(id uint) (*entity.User_Auth, error) {
	auth := entity.User_Auth{}
	has, err := db.Get(&auth)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &auth, nil
}

func NewUserCore() UserCore {
	return &userCore{}
}
