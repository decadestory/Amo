package ammodel

import (
	"time"
)

type AmConfigExcel struct {
	ID         int       `gorm:"primaryKey"`
	ConfigCode string    `gorm:"column:ConfigCode"`
	ColKey     string    `gorm:"column:ColKey"`
	ColValue   string    `gorm:"column:ColValue"`
	LineID     string    `gorm:"column:LineId"`
	DomainID   string    `gorm:"column:DomainId"`
	AddTime    time.Time `gorm:"column:AddTime"`
	IsValid    bool      `gorm:"column:IsValid"`
}

// TableName 会将 AmConfigExcel 的表名重写为 `AmConfigExcel`
func (AmConfigExcel) TableName() string {
	return "AmConfigExcel"
}
