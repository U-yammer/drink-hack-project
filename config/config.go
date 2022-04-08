package config

// アプリ全体の設定に関連するファイル

import (
	"drink_hack_project/utils"
	"gopkg.in/go-ini/ini.v1"
	"log"
	_ "log"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

// 最初に実行される
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {log.Fatalln(err)}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
