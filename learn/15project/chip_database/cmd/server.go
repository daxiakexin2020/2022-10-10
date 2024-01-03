package main

import (
	"chip_database/conf"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	run()
}

func run() {

	k, err := initApp(gin.New(), conf.GetWebServerConfig(), conf.GetSqliteDatabaseConfig(), conf.GetFileHandleProxyOnce())
	if err != nil {
		os.Exit(-1)
	}
	k.Run()
}