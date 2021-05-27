package ammodel

import "time"

type LogAmBus struct {
	ID       int       `gorm:"primaryKey"`
	SrcGId   int       `gorm:"column:SrcGId"`
	SrcId    int       `gorm:"column:SrcId"`
	LogType  string    `gorm:"column:LogType"`
	LogLevel int       `gorm:"column:LogLevel"`
	LogPath  string    `gorm:"column:LogPath"`
	LogInfo  string    `gorm:"column:LogInfo"`
	ExtInfo1 string    `gorm:"column:ExtInfo1"`
	ExtInfo2 string    `gorm:"column:ExtInfo2"`
	AddTime  time.Time `gorm:"column:AddTime"`
}

func (LogAmBus) TableName() string {
	return "LogAmBus"
}
