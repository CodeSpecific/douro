package config

import (
	"github.com/go-ini/ini"
)

var Conf Config = newConfig()

type Config interface {
	GetDbConnectStr() string
}

type iniConfig struct {
	f            *ini.File
	dbConnectStr string
}

func (c *iniConfig) GetDbConnectStr() string {
	return c.dbConnectStr
}

func (c *iniConfig) init() *iniConfig {
	var err error
	c.f, err = ini.Load("config/douro.ini")
	if err != nil {
		panic(err)
	}
	c.dbConnectStr = c.f.Section("mysql").Key("connectStr").String()
	return c
}

func newConfig() Config {
	return new(iniConfig).init()
}
