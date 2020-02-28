package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// 構造体ConfigList定義
type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

// グローバル変数Config定義
var Config ConfigList

func init() {
	// config.iniを読み込み
	cfg, err := ini.Load("config.ini")
	// config.iniが読み込めなかった時の処理
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		// エラーコード1で処理を抜ける。
		os.Exit(1)
	}

	// グローバル変数Configにconfig.iniで取得した値を代入する。
	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:   cfg.Section("gotrading").Key("log_file").String(),
	}
}
