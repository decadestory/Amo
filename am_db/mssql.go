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
func GetMssqlConfigMapper() []ammodel.ConfigMapper {

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	var mappers []ammodel.ConfigMapper
	db.Limit(10).Order("id asc").Find(&mappers)

	return mappers
}
