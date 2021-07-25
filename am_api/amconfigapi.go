package amapi

import (
	"fmt"
	"net/http"

	amconfig "atom_micro/am_config"
	amlib "atom_micro/am_lib"
	ammodel "atom_micro/am_model"

	"github.com/gin-gonic/gin"
)

func StartConfigApi() {
	engin := gin.Default()

	engin.Use(amlib.Recover)

	engin.GET("/ping", ping_config)
	engin.POST("/set", SetConfig)
	engin.POST("/get/:code", GetConfig)
	//add more
	go http.ListenAndServe(":2903", engin)
}

func ping_config(c *gin.Context) {
	c.JSON(0, gin.H{"message": "pong"})
}

func SetConfig(c *gin.Context) {
	var data ammodel.AmConfigModel
	err := c.ShouldBind(&data)

	fmt.Println("setConfig -- ShouldBind", data)

	if err != nil {
		fmt.Println("error:", err)
		c.JSON(0, ammodel.Error(err.Error()))
		return
	}

	amconfig.SetConfig(data)
	c.JSON(0, ammodel.Ok(true))

}

func GetConfig(c *gin.Context) {
	code := c.Param("code")
	res := amconfig.GetConfig(code)
	if res.ID == 0 {
		c.JSON(0, ammodel.Error("code不存在"))
	}
	c.JSON(0, ammodel.Ok(res))
}
