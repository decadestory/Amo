package amconfig

import (
	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"
	"fmt"
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
	ent.StartTime = time.Time(model.StartTime)
	ent.EndTime = time.Time(model.EndTime)

	fmt.Println("setConfig -- NewConverter", ent)

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

	model.AddTime = ammodel.Time(ent.AddTime)
	model.StartTime = ammodel.Time(ent.StartTime)
	model.EndTime = ammodel.Time(ent.EndTime)

	return model
}
