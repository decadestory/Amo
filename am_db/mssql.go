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

// 获取映射
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

// 添加接口日志
func AddApiLog(lg ammodel.LogAmInterface) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Create(&lg)
}

// 添加业务日志
func AddBusLog(lg ammodel.LogAmBus) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Create(&lg)
}

// 添加错误日志
func AddErrorLog(lg ammodel.LogAmError) {
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Create(&lg)
}
