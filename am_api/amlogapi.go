package amapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	logger "atom_micro/am_log"
	ammodel "atom_micro/am_model"
)

func StartLogApi() {
	engin := gin.Default()
	engin.POST("/", ping_log)
	engin.POST("/logApi", serveHTTPApi)
	engin.POST("/logBus", serveHTTPBus)
	engin.POST("/logError", serveHTTPError)
	//add more
	go http.ListenAndServe(":2902", engin)
}

func ping_log(c *gin.Context) {
	c.JSON(200, gin.H{"message": "日志服务启动成功！！"})
}

func serveHTTPApi(c *gin.Context) {
	var data ammodel.LogAmInterface
	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(200, ammodel.Error(err.Error()))
		return
	}

	logger.LogApi(data.LogPath, data.Parameter, data.ExecuteTime)
	fmt.Println("ok:", data)

	c.JSON(200, ammodel.Ok(data))
}

func serveHTTPBus(c *gin.Context) {
	var data ammodel.LogAmBus
	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(200, ammodel.Error(err.Error()))
		return
	}

	logger.LogBus(data.SrcGId, data.SrcId, data.LogLevel, data.LogType, data.LogPath, data.LogInfo, data.ExtInfo1, data.ExtInfo2)
	fmt.Println("ok:", data)

	c.JSON(200, ammodel.Ok(data))
}

func serveHTTPError(c *gin.Context) {
	var data ammodel.LogAmError
	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(200, ammodel.Error(err.Error()))
		return
	}

	logger.LogError(data.SrcGId, data.SrcId, data.LogLevel, data.LogType, data.LogPath, data.LogInfo, data.ExtInfo1, data.ExtInfo2)
	fmt.Println("ok:", data)

	c.JSON(200, ammodel.Ok(data))
}
