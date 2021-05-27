package ammodel

import "time"

type LogAmInterface struct {
	ID          int       `gorm:"primaryKey"`
	SrcGId      int       `gorm:"column:SrcGId"`
	SrcId       int       `gorm:"column:SrcId"`
	LogType     string    `gorm:"column:LogType"`
	LogPath     string    `gorm:"column:LogPath"`
	Parameter   string    `gorm:"column:Parameter"`
	ExecuteTime int       `gorm:"column:ExecuteTime"`
	AddTime     time.Time `gorm:"column:AddTime"`
}

func (LogAmInterface) TableName() string {
	return "LogAmInterface"
}
