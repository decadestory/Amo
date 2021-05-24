package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	amapi "atom_micro/am_api"
)

// 实际中应该用更好的变量名
var (
	run  string
	info bool
	v    bool
	h    bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version")
	flag.BoolVar(&info, "info", false, "show amo component details")
	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&run, "run", "", "send `signal` to a master process: log, config, gateway")
	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	fmt.Println("amo is ready")
	flag.Parse()

	if h {
		flag.Usage()
	} else if run != "" {
		selServsToStart(run)
	} else {
		flag.Usage()
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, "amo version: 1.0.0 \n Usage: amo [-hv] [-info] [-run componentsname] \n Options:\n")
	flag.PrintDefaults()
}

func selServsToStart(filterStr string) {
	ss := strings.Split(filterStr, ",")

	for k, v := range ss {
		fmt.Println(k)

		switch v {
		case "gateway":
			amapi.StartGateWayAPI()
		case "log":
			amapi.StartLogApi()
		case "config":
			amapi.StartConfigApi()
		case "auth":
			amapi.StartAuthApi()
		default:
			fmt.Println("服务类型指定有误：", v)
		}
	}

	fmt.Println("启动成功")
	//阻塞程序
	select {}

}
