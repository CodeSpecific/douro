package model

import (
	"time"

	"github.com/CodeSpecific/douro/service"
)

// UserViewModel (ViewModel)用户信息
type UserViewModel struct {
	ID           uint
	Name         string
	Phone        string
	Gender       service.Gender
	Avatar       string
	RegisterMode service.RegisterMode
	CreateTime   time.Time
}
