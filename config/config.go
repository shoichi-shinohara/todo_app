package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      config.Section("web").Key("port").MustString("8080"),
		SQLDriver: config.Section("db").Key("driver").String(),
		DbName:    config.Section("db").Key("name").String(),
		LogFile:   config.Section("web").Key("logfile").String(),
		Static:    config.Section("web").Key("static").String(),
	}
}
