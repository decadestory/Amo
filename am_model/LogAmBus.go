package ammodel

import "time"

type LogAmBus struct {
	ID       int `gorm:"primaryKey"`
	SrcGId   int
	SrcId    int
	LogType  string
	LogLevel int
	LogPath  string
	LogInfo  string
	ExtInfo1 string
	ExtInfo2 string
	AddTime  time.Time `gorm:"column:AddTime"`
}

func (LogAmBus) TableName() string {
	return "LogAmBus"
}
