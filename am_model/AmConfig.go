package ammodel

import (
	"time"
)

type AmConfig struct {
	ID          int       `gorm:"primaryKey"`
	ConfigCode  string    `gorm:"column:ConfigCode"`
	ConfigName  string    `gorm:"column:ConfigName"`
	ConfigDesc  string    `gorm:"column:ConfigDesc"`
	ConfigValue string    `gorm:"column:ConfigValue"`
	ExtValue1   string    `gorm:"column:ExtValue1"`
	ExtValue2   string    `gorm:"column:ExtValue2"`
	StartTime   time.Time `gorm:"column:StartTime" json:"start_at" time_format:"2006-01-02" time_utc:"1"`
	EndTime     time.Time `gorm:"column:EndTime"`
	HasDetail   bool      `gorm:"column:HasDetail"`
	ParentCode  string    `gorm:"column:ParentCode"`
	DomainID    string    `gorm:"column:DomainId"`
	AddTime     time.Time `gorm:"column:AddTime"`
	IsValid     bool      `gorm:"column:IsValid"`
}

// TableName 会将 AmConfig 的表名重写为 `AmConfig`
func (AmConfig) TableName() string {
	return "AmConfig"
}
