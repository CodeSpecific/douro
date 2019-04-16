package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

var Conf Config = newConfig()

type Config interface {
	GetDbConnectStr() string
}

type iniConfig struct {
	f            *ini.File
	DbConnectStr string
}

func (c *iniConfig) GetDbConnectStr() string {
	return c.f.Section("mysql").Key("connectStr").String()
}

func (c *iniConfig) init() *iniConfig {
	var err error
	c.f, err = ini.Load("config/douro.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return c
}

func newConfig() Config {
	return new(iniConfig).init()
}
