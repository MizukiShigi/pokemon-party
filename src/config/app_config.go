package config

import (
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	AccessLog string
	SystemLog string
}

var Config ConfigList

func init() {
	loadConfig()
}

func loadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		AccessLog: cfg.Section("web").Key("accesslog").MustString("log/app/access.log"),
		SystemLog: cfg.Section("web").Key("systemlog").MustString("log/app/system.log"),
	}
}
