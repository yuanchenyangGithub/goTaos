package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// r.Use(Cors())
	r.GET("/export", func(c *gin.Context) {
		url := "https://monitor.genyon.cn/genyon-admin/exportGb32960Data"
		method := "POST"

		vehicleNo := c.Query("vehicleNo")
		clientId := c.Query("clientId")
		// vehicleNo := c.PostForm("vehicleNo")
		fmt.Println(vehicleNo)
		// clientId := c.PostForm("clientId")
		fmt.Println(clientId)
		time := c.Query("time")
		fmt.Println(time)

		token := c.Query("token")

		// time := "2021-12-08"
		strArrayVehicleNo := [...]string{vehicleNo}
		strArrayClientId := [...]string{clientId}

		for i := 0; i < len(strArrayVehicleNo); i++ {

			var strNewReader string = "{\"vehicleNo\":\"" + strArrayVehicleNo[i] + "\",\"clientId\":\"" + strArrayClientId[i] + "\",\"time\":\"" + time + "\"}`"
			fmt.Println(strNewReader)
			payload := strings.NewReader(strNewReader)

			// payload := strings.NewReader(`{"vehicleNo":"皖N01888D","clientId":"LB9HFHBN1GAJHW592","time":"2021-12-08"}`)

			client := &http.Client{}
			req, err := http.NewRequest(method, url, payload)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("token", token)
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println(string(body))

			c.Writer.Header().Add("Content-Type", "application/octet-stream")
			c.Writer.Header().Add("Content-Disposition", "attachment; filename=\""+strArrayVehicleNo[i]+"-"+time+".xlsx"+"\"")
			c.Writer.Write(body)
		}

	})
	_ = r.Run()
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
