package service

import (
	"time"

	"github.com/CodeSpecific/douro/core"
)

// Gender 性别枚举
type Gender int

const (
	Unknown Gender = iota
	Male
	Female
)

// RegisterMode 注册渠道
type RegisterMode int

const (
	Phone RegisterMode = iota
	WeiXin
)

// User (BO)用户业务对象
type User struct {
	ID           uint
	Name         string
	Phone        string
	Gender       Gender
	Avatar       string
	RegisterMode RegisterMode
	ThirdPartyID string
	EncryptPwd   string
	CreateTime   time.Time
}

type UserService interface {
	GetUser(id uint) (*User, error)
}

type userService struct {
	core.UserCore
}

func (s *userService) GetUser(id uint) (*User, error) {
	profile := s.GetUserProfile(id)
	auth := s.GetUserAuth(id)
	user := convertToUser(profile, auth)
	return user, nil
}

func NewUserService() UserService {
	return &userService{core.NewUserCore()}
}
