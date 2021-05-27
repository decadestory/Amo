package ammodel

import "time"

type LogAmError struct {
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

func (LogAmError) TableName() string {
	return "LogAmError"
}
