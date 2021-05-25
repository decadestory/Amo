package amapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	amdb "atom_micro/am_db"
	ammodel "atom_micro/am_model"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var gcc *cache.Cache = cache.New(5*time.Minute, 5*time.Minute)

// StartGateWayAPI 启动服务
func StartGateWayAPI() {
	engin := gin.Default()
	engin.GET("/*url", serveHTTP)
	engin.POST("/*url", serveHTTP)

	go http.ListenAndServe(":2901", engin)
}

func serveHTTP(c *gin.Context) {

	path := c.Request.URL.Path
	fmt.Println("请求路径", path)

	var mps []ammodel.ConfigMapper
	foo, found := gcc.Get("mssql-mappers")
	if found {
		fmt.Println("found:", foo.(string))
		json.Unmarshal([]byte(foo.(string)), &mps)
	} else {
		mps = amdb.GetMssqlConfigMapper()
		mpsJson, _ := json.Marshal(mps)
		gcc.Set("mssql-mappers", string(mpsJson), cache.DefaultExpiration)
		fmt.Println("not found:", string(mpsJson))
	}

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
