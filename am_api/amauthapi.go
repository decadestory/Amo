package amapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartAuthApi() {
	engin := gin.Default()
	engin.GET("/ping", ping_auth)
	//add more
	go http.ListenAndServe(":2904", engin)
}

func ping_auth(c *gin.Context) {
	c.JSON(0, gin.H{"message": "pong"})
}
