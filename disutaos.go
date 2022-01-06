package main

import (
	logger "disutaos/logger"
	routers "disutaos/routers"
	"fmt"

	taos "disutaos/taos"

	taosconsumer "disutaos/taosconsume"

	"github.com/gin-gonic/gin"
	loggerhub "github.com/wonderivan/logger"
)

func main() {

	// Force log's color
	gin.ForceConsoleColor()

	go taosconsumer.Consumer()

	r := gin.Default()

	err := taos.InitTaoSql()
	if err != nil {
		fmt.Println("sartup service failed, err:%v\n")
	}

	defer taos.Close()

	// 3.监听端口，默在800
	// Run("里面不指定端口号默认为800"
	r.Use(logger.LogerMiddleware())
	routers.SetRouter(r)

	if err := r.Run(":10010"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

	loggerhub.Debug("kaishi")

	// defer taos.SqlDB.Close()
}
