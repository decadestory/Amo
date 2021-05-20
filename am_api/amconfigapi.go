package amapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartConfigApi() {
	engin := gin.Default()
	engin.GET("/ping", ping_config)
	//add more
	go http.ListenAndServe(":2903", engin)
}

func ping_config(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
