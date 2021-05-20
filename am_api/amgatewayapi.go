package amapi

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func StartGateWayApi() {
	engin := gin.Default()
	engin.GET("/*url", serveHttp)
	engin.POST("/*url", serveHttp)

	//add more
	go http.ListenAndServe(":2901", engin)
}

func serveHttp(c *gin.Context) {
	target := "sfa-cn.lorealchina.com"

	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = target
		req.URL.Path = "/api/CPDTest/GetQrPageTxt"
		// add custom headers
		req.Header["my-header"] = []string{req.Header.Get("my-header")}
		// delete Origin headers
		delete(req.Header, "My-Header")
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
