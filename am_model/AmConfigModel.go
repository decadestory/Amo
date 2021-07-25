package ammodel

type AmConfigModel struct {
	ID          int    `gorm:"primaryKey"`
	ConfigCode  string `gorm:"column:ConfigCode"`
	ConfigName  string `gorm:"column:ConfigName"`
	ConfigDesc  string `gorm:"column:ConfigDesc"`
	ConfigValue string `gorm:"column:ConfigValue"`
	ExtValue1   string `gorm:"column:ExtValue1"`
	ExtValue2   string `gorm:"column:ExtValue2"`
	StartTime   Time   `gorm:"column:StartTime" time_format:"2006-01-02 15:04:05"`
	EndTime     Time   `gorm:"column:EndTime" time_format:"2006-01-02 15:04:05"`
	HasDetail   bool   `gorm:"column:HasDetail"`
	ParentCode  string `gorm:"column:ParentCode"`
	DomainID    string `gorm:"column:DomainId"`
	AddTime     Time   `gorm:"column:AddTime" time_format:"2006-01-02 15:04:05"`
	IsValid     bool   `gorm:"column:IsValid"`
}
