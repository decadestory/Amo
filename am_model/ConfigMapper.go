package ammodel

import (
	"time"
)

type ConfigMapper struct {
	ID              int `gorm:"primaryKey"`
	ServiceId       string
	ServiceName     string
	UpSteamHost     string
	UpSteamPath     string
	DownSteamScheme string
	DownSteamHost   string
	DownSteamPort   int
	DownSteamPath   string
	NeedAuth        bool
	AddTime         time.Time
	Enable          bool
}

// TableName 会将 ConfigMapper 的表名重写为 `ConfigMapper`
func (ConfigMapper) TableName() string {
	return "ConfigMapper"
}
