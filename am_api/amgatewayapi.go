package amapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	amdb "atom_micro/am_db"
	logger "atom_micro/am_log"
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
	var elapsed int
	path := c.Request.URL.Path
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点：读取完body要重新赋值上
	defer logger.LogApi(path, string(data), int(elapsed))

	//缓存
	var mps []ammodel.AmProxyMapper
	foo, found := gcc.Get("proxy-mappers")
	if found {
		fmt.Println("found!!")
		json.Unmarshal([]byte(foo.(string)), &mps)
	} else {
		mps = amdb.GetMssqlAmConfigMapper()
		mpsJson, _ := json.Marshal(mps)
		gcc.Set("proxy-mappers", string(mpsJson), cache.DefaultExpiration)
		fmt.Println("not found!!")
	}

	//查询映射
	var mapped ammodel.AmProxyMapper
	for _, v := range mps {
		if strings.HasPrefix(path, v.UpSteamPath) {
			mapped = v
		}
	}

	//代理转发
	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", mapped.DownSteamHost)
		req.URL.Scheme = mapped.DownSteamScheme
		req.URL.Host = mapped.DownSteamHost
		req.URL.Path = path
		// add custom headers
		// req.Header["my-header"] = []string{req.Header.Get("my-header")}
		// delete Origin headers
		// delete(req.Header, "My-Header")
	}

	//记录时间
	start := time.Now() // 获取当前时间
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)

	elapsed = (int)(time.Since(start) / 1000000)
	fmt.Println("execute time:", elapsed, "ms")
}
