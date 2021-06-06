package ammodel

import (
	"database/sql"
	"time"
)

type AmConfigDetail struct {
	ID             int          `gorm:"primaryKey"`
	RelID          string       `gorm:"column:RelId"`
	ConfigCode     string       `gorm:"column:ConfigCode"`
	ConfigValue    string       `gorm:"column:ConfigValue"`
	ConfigType     string       `gorm:"column:ConfigType"`
	DimensionCode  string       `gorm:"column:DimensionCode"`
	DimensionCode2 string       `gorm:"column:DimensionCode2"`
	ExtValue1      string       `gorm:"column:ExtValue1"`
	ExtValue2      string       `gorm:"column:ExtValue2"`
	StartTime      sql.NullTime `gorm:"column:StartTime"`
	EndTime        sql.NullTime `gorm:"column:EndTime"`
	DomainID       string       `gorm:"column:DomainId"`
	AddTime        time.Time    `gorm:"column:AddTime"`
	IsValid        bool         `gorm:"column:IsValid"`
}

// TableName 会将 AmConfigDetail 的表名重写为 `AmConfigDetail`
func (AmConfigDetail) TableName() string {
	return "AmConfigDetail"
}
