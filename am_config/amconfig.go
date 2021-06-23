package amconfig

import (
	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"
	"time"

	"github.com/petersunbag/coven"
)

// SetConfig 添加修改配置
func SetConfig(model ammodel.AmConfigModel) bool {

	//model to entity
	var c, err = coven.NewConverter(ammodel.AmConfig{}, ammodel.AmConfigModel{})
	if err != nil {
		panic(err)
	}

	ent := ammodel.AmConfig{}
	err = c.Convert(&ent, &model)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	ent.AddTime = now
	ent.StartTime = now
	ent.EndTime = now.AddDate(5000, 0, 0)

	amdb.SetConfig(ent)
	return true
}

// GetConfig 添加修改配置
func GetConfig(code string) ammodel.AmConfigModel {
	//model to entity
	var c, err = coven.NewConverter(ammodel.AmConfigModel{}, ammodel.AmConfig{})
	if err != nil {
		panic(err)
	}

	ent := amdb.GetConfig(code)
	model := ammodel.AmConfigModel{}
	err = c.Convert(&model, &ent)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	model.AddTime = ammodel.Time(now)
	model.StartTime = ammodel.Time(now)
	model.EndTime = ammodel.Time(now.AddDate(5000, 0, 0))

	return model
}
