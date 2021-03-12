package config

import (
	"time"

	"github.com/caarlos0/env"
)

// ApplicationsName 機能名
const ApplicationsName = "ichi-infra-challenge"

// 設定値
var (
	AppVersion      string
	AppStage        string
	LogLevel        string
	TimeZone        = time.UTC
	MySQLConnection string
)

type config struct { // nolint:maligned
	Version         string `env:"APP_VERSION" envDefault:""`
	AppStage        string `env:"APP_STAGE" envDefault:"local"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"warn"`
	TimeZone        string `env:"TIME_ZONE" envDefault:"UTC"`
	MySQLConnection string `env:"MYSQL_CONNECTION" envDefault:"local:local@tcp(127.0.0.1:3306)/infra-challenge?charset=utf8&parseTime=true"`
}

func init() {
	Set()
}

// Set sets configurations via envoronment variables
func Set() {
	cfg := config{}
	env.Parse(&cfg)
	AppVersion = cfg.Version
	AppStage = cfg.AppStage
	LogLevel = cfg.LogLevel
	if location, err := time.LoadLocation(cfg.TimeZone); err == nil {
		TimeZone = location
	}
	MySQLConnection = cfg.MySQLConnection

}
