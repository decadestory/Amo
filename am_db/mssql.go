package amdb

import (
	"fmt"
	"log"

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
	if amc.ID == 0 {
		fmt.Println("not find -- create")
		db.Create(&cm)
	} else {
		fmt.Println("find -- update")
		db.Model(&cm).Where("ID=?", amc.ID).Updates(cm)
	}
}

// GetConfig 获取主配置
func GetConfig(code string) ammodel.AmConfig {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	var amc ammodel.AmConfig
	db.First(&amc, "ConfigCode=?", code)
	return amc
}
