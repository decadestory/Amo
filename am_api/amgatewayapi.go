package amapi

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"

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

	path := c.Request.URL.Path
	fmt.Println("请求路径", path)

	mps := amdb.GetMssqlConfigMapper()
	log.Println("长度：", len(mps))

	var mapped ammodel.ConfigMapper
	for _, v := range mps {
		if strings.HasPrefix(path, v.UpSteamPath) {
			mapped = v
		}
	}

	fmt.Println(mapped)

	director := func(req *http.Request) {
		req.URL.Scheme = mapped.DownSteamScheme
		req.URL.Host = mapped.DownSteamHost
		req.URL.Path = path
		// add custom headers
		// req.Header["my-header"] = []string{req.Header.Get("my-header")}
		// delete Origin headers
		// delete(req.Header, "My-Header")
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
