package amdb

import (
	"fmt"
	"log"
	"time"

	ammodel "atom_micro/am_model"

	"github.com/Unknwon/goconfig"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var connString string

func init() {

	c, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatal("Config load faild:", err.Error())
	}

	server, _ := c.GetValue("dbconf", "server")
	user, _ := c.GetValue("dbconf", "user")
	pwd, _ := c.GetValue("dbconf", "pwd")
	// port, _ := c.GetValue("dbconf", "port")
	database, _ := c.GetValue("dbconf", "database")

	connString = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, pwd, server, database)
	log.Println(connString)
}

// GetMssqlAmConfigMapper 获取映射
func GetMssqlAmConfigMapper() []ammodel.AmProxyMapper {

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	var mappers []ammodel.AmProxyMapper
	db.Limit(10).Order("id asc").Find(&mappers)

	log.Println("获取", len(mappers), "条数据")
	return mappers
}

// AddApiLog 添加接口日志
func AddApiLog(lg ammodel.LogAmInterface) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Create(&lg)
}

// AddBusLog 添加业务日志
func AddBusLog(lg ammodel.LogAmBus) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Create(&lg)
}

// AddErrorLog 添加错误日志
func AddErrorLog(lg ammodel.LogAmError) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Create(&lg)
}

// SetConfig 添加或修改主配置
func SetConfig(cm ammodel.AmConfig) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	var amc ammodel.AmConfig
	db.First(&amc, "ConfigCode=?", cm.ConfigCode)
	fmt.Println("amc -->", amc)
	now := time.Now()
	if amc.ID == 0 {
		cm.AddTime = now
		cm.StartTime = now
		cm.EndTime = now.AddDate(5000, 0, 0)
		cm.IsValid = true
		fmt.Println("not find -- create")
		db.Create(&cm)
	} else {
		fmt.Println("find -- update", cm)
		cm.AddTime = now
		umap := map[string]interface{}{
			"ConfigName":  cm.ConfigName,
			"ConfigDesc":  cm.ConfigDesc,
			"ConfigValue": cm.ConfigValue,
			"ExtValue1":   cm.ExtValue1,
			"ExtValue2":   cm.ExtValue2,
			"HasDetail":   cm.HasDetail,
			"ParentCode":  cm.ParentCode,
			"DomainID":    cm.DomainID,
			"AddTime":     cm.AddTime,
			"IsValid":     cm.IsValid,
		}
		if !cm.StartTime.IsZero() {
			fmt.Println("!StartTime.IsZero")
			umap["StartTime"] = cm.StartTime
		}

		if !cm.EndTime.IsZero() {
			fmt.Println("!EndTime.IsZero")
			umap["EndTime"] = cm.EndTime
		}

		db.Model(&cm).Where("ID=?", amc.ID).Updates(umap)

	}
}

// GetConfig 获取主配置
func GetConfig(code string) ammodel.AmConfig {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	var amc ammodel.AmConfig
	db.First(&amc, "ConfigCode=? and IsValid=1 and getdate() BETWEEN StartTime AND EndTime ", code)
	fmt.Println("GetConfig -- ", amc)
	if amc.ID == 0 {
		panic("配置不存在")
	}

	return amc
}
