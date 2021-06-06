package amconfig

import (
	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"
	"time"
)

// SetConfig 添加修改配置
func SetConfig(model ammodel.AmConfig) bool {

	model.AddTime = time.Now()
	model.IsValid = true

	amdb.SetConfig(model)
	return true
}

// GetConfig 添加修改配置
func GetConfig(code string) ammodel.AmConfig {
	return amdb.GetConfig(code)
}
