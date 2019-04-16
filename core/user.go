package core

import "github.com/CodeSpecific/douro/core/entity"

// UserCore DAO对象
type UserCore interface {
	GetUserProfile(id uint) (*entity.User_Profile, error)
	GetUserAuth(id uint) (*entity.User_Auth, error)
}

type userCore struct {
}

func (c *userCore) GetUserProfile(id uint) (*entity.User_Profile, error) {
	user := entity.User_Profile{}
	_, err := db.Get(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *userCore) GetUserAuth(id uint) (*entity.User_Auth, error) {
	return &entity.User_Auth{}, nil
}

func NewUserCore() UserCore {
	return &userCore{}
}
