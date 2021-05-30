package ammodel

import (
	"time"
)

type AmProxyMapper struct {
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

// TableName 会将 AmProxyMapper 的表名重写为 `AmProxyMapper`
func (AmProxyMapper) TableName() string {
	return "AmProxyMapper"
}
