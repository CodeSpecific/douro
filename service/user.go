package service

import (
	"time"

	"github.com/CodeSpecific/douro/core"
)

// Gender 性别枚举
type Gender uint

const (
	保密 Gender = iota
	男
	女
)

// RegisterMode 注册渠道
type RegisterMode uint

const (
	电话 RegisterMode = iota
	微信
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
