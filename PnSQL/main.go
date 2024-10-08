package main

import (
	"flag"
	"github.com/xiaohongshu/PnSql/server/initialize"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "", "port to use")
	flag.Parse()
	args := map[string]string{"port": port}
	initialize.InitConfig()
	initialize.InitRouter(args)

}
