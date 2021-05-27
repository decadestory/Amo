package ammodel

import (
	"time"
)

type AmConfigMapper struct {
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

// TableName 会将 AmConfigMapper 的表名重写为 `AmConfigMapper`
func (AmConfigMapper) TableName() string {
	return "AmConfigMapper"
}
