package taoscontroller

import (

	// taos "disutaos/taos"
	tbatter "disutaos/modules"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taoscontrollers struct {
}

func PostDevHistoryInfo(c *gin.Context) {

	clientId, err := strconv.Atoi(c.Query("clientId"))
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	tBatteryHistoryDataList, err := tbatter.FindListTBatterHD(clientId, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "SUCCESS",
		"data":    tBatteryHistoryDataList,
		"retain":  0,
	})

	// if tbatterList, err := tbatter.Hello(); err != nil {
	// 	c.JSON(200, gin.H{"error": err.Error()})
	// } else {
	// 	c.JSON(200, tbatterList)
	// }
}

func GetTotalLoopNumber(c *gin.Context) {
	numLtime := tbatter.GetTotalLoopNumber()
	c.JSON(200, numLtime)
}

func PostDevNewHistoryInfo(c *gin.Context) {

	clientId, err := strconv.Atoi(c.Query("clientId"))
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	tBatteryLocationDataList, err := tbatter.PostDevNewHistoryInfo(clientId, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "SUCCESS",
		"data":    tBatteryLocationDataList,
		"retain":  0,
	})
}

func PostBatteryHistoryInfo(c *gin.Context) {

	clientId, err := strconv.Atoi(c.Query("clientId"))
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	tBatteryLocationDataList, err := tbatter.PostBatteryHistoryInfo(clientId, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "SUCCESS",
		"data":    tBatteryLocationDataList,
	})

}
