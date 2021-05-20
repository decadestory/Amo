package amapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartLogApi() {
	engin := gin.Default()
	engin.GET("/ping", ping_log)
	//add more
	go http.ListenAndServe(":2902", engin)
}

func ping_log(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
