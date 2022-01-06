package routers

import (
	"fmt"
	"net/http"

	taoscontrollers "disutaos/controllers"

	authorize "disutaos/authorize"
	cross "disutaos/cross"

	"github.com/gin-gonic/gin"
)

func SetRouter(e *gin.Engine) {

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	//定义路由的GET方法及响应的处理函数
	e.Use(cross.CrossMiddleware())

	http.HandleFunc("/", index)
	e.GET("/hello", Hello)

	e.Use(authorize.Authorize())
	// 电池异常告警数据查询
	e.POST("/postDevHistoryInfo", taoscontrollers.PostDevHistoryInfo)
	// 大屏展示显示
	e.GET("/getTotalLoopNumber", taoscontrollers.GetTotalLoopNumber)
	// 调去历史设备轨迹信息
	e.POST("/postDevNewHistoryInfo", taoscontrollers.PostDevNewHistoryInfo)
	// 调取历史电池所有信息
	e.POST("/postBatteryHistoryInfo", taoscontrollers.PostBatteryHistoryInfo)

	e.GET("/hello2", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World two")
	})
}

func Hello(c *gin.Context) {
	fmt.Println("hello world！")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Gin",
	})
}

//首页跳转
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}

}
