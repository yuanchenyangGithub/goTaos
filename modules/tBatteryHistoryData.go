package tbatter

import (
	taos "disutaos/taos"
	"fmt"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type TBatteryHistoryData struct {
	Ts                 string  `json:"ts" form:"ts"`
	Altitude           int     `json:"altitude" form:"altitude"`
	Latitude           float64 `json:"latitude" form:"latitude"`
	Longitude          string  `json:"longitude" form:"longitude"`
	Mobile             int     `json:"mobile" form:"mobile"`
	InfoTime           string  `json:"infoTime" form:"infoTime"`
	Speed              int     `json:"speed" form:"speed"`
	Direction          int     `json:"direction" form:"direction"`
	BatteryId          string  `json:"batteryId" form:"batteryId"`
	TotalVoltage       float32 `json:"totalVoltage" form:"totalVoltage"`
	TotalCurrent       float32 `json:"totalCurrent" form:"totalCurrent"`
	Soc                int     `json:"soc" form:"soc"`
	CellQuantity       int     `json:"cellQuantity" form:"cellQuantity"`
	CellVoltageDetail  string  `json:"cellVoltageDetail" form:"cellVoltageDetail"`
	TempQuantity       int     `json:"tempQuantity" form:"tempQuantity"`
	TempDetailInfo     string  `json:"tempDetailInfo" form:"tempDetailInfo"`
	BMSTemp            float32 `json:"bMSTemp" form:"bMSTemp"`
	ResidualCapacity   float32 `json:"residualCapacity" form:"residualCapacity"`
	CurrentCapacity    float32 `json:"currentCapacity" form:"currentCapacity"`
	LoopTimes          int     `json:"loopTimes" form:"loopTimes"`
	LifeSignal         float32 `json:"lifeSignal" form:"lifeSignal"`
	GsmSignalIntensity string  `json:"gsmSignalIntensity" form:"gsmSignalIntensity"`
	LocationSatNum     int     `json:"locationSatNum" form:"locationSatNum"`
}

/*
 * @Author: chenyang_yuan
 * @date: 2021年11月30日
 * @content：根据查询电池
 */
func FindListTBatterHD(clientId int, startTime string, endTime string) (tBatteryHistoryDataList []TBatteryHistoryData, err error) {
	tBatteryHistoryDataList = make([]TBatteryHistoryData, 0)

	sqlStr := "select ts,altitude,latitude,longitude,mobile,info_time,speed,direction,battery_id,total_voltage,total_current,soc,cell_quantity,cell_voltage_detail,temp_quantity,temp_detail_info,b_m_s_temp,residual_capacity,current_capacity,loop_times, life_signal, gsm_signal_intensity,location_sat_num  from devDB.t_battery_history_data where mobile = " + strconv.Itoa(clientId) + " and info_time > " + startTime + " and info_time < " + endTime + " order by ts asc"
	fmt.Println(sqlStr)
	taos.CheckErr(err, sqlStr)
	res, err := taos.SqlDB.Query(sqlStr)
	defer res.Close()

	if err != nil {
		return
	}
	for res.Next() {
		var tBatteryHistoryData TBatteryHistoryData
		// Scan 方法会从输入端读取数据并将处理结果存入接收端，接收端必须是有效的指针。
		res.Scan(&tBatteryHistoryData.Ts,
			&tBatteryHistoryData.Altitude,
			&tBatteryHistoryData.Latitude,
			&tBatteryHistoryData.Longitude,
			&tBatteryHistoryData.Mobile,
			&tBatteryHistoryData.InfoTime,
			&tBatteryHistoryData.Speed,
			&tBatteryHistoryData.Direction,
			&tBatteryHistoryData.BatteryId,
			&tBatteryHistoryData.TotalVoltage,
			&tBatteryHistoryData.TotalCurrent,
			&tBatteryHistoryData.Soc,
			&tBatteryHistoryData.CellQuantity,
			&tBatteryHistoryData.CellVoltageDetail,
			&tBatteryHistoryData.TempQuantity,
			&tBatteryHistoryData.TempDetailInfo,
			&tBatteryHistoryData.BMSTemp,
			&tBatteryHistoryData.ResidualCapacity,
			&tBatteryHistoryData.CurrentCapacity,
			&tBatteryHistoryData.LoopTimes,
			&tBatteryHistoryData.LifeSignal,
			&tBatteryHistoryData.GsmSignalIntensity,
			&tBatteryHistoryData.LocationSatNum)
		tBatteryHistoryDataList = append(tBatteryHistoryDataList, tBatteryHistoryData)
	}
	return
}

func GetTotalLoopNumber() uint32 {
	var NumLtime uint32
	sqlStr := "select sum(ltime) NumLtime from ( SELECT max(loop_times) ltime ,mobile  from devDB.t_battery_history_data group by mobile)"
	err := taos.SqlDB.QueryRow(sqlStr).Scan(&NumLtime)
	taos.CheckErr(err, sqlStr)
	if err != nil {
		fmt.Println("查询出错")
	}
	return NumLtime
}

func PostBatteryHistoryInfo(clientId int, startTime string, endTime string) (tBatteryHistoryDataList []TBatteryHistoryData, err error) {
	tBatteryHistoryDataList = make([]TBatteryHistoryData, 0)

	sqlStr := "select ts,altitude,latitude,longitude,mobile,info_time,speed,direction,battery_id,total_voltage,total_current,soc,cell_quantity,cell_voltage_detail,temp_quantity,temp_detail_info,b_m_s_temp,residual_capacity,current_capacity,loop_times,gsm_signal_intensity,location_sat_num  from devDB.t_battery_history_data where mobile = " + strconv.Itoa(clientId) + " and info_time > " + startTime + " and info_time < " + endTime + " order by ts asc"
	fmt.Println(sqlStr)
	taos.CheckErr(err, sqlStr)
	res, err := taos.SqlDB.Query(sqlStr)
	defer res.Close()

	if err != nil {
		return
	}
	for res.Next() {
		var tBatteryHistoryData TBatteryHistoryData
		// Scan 方法会从输入端读取数据并将处理结果存入接收端，接收端必须是有效的指针。
		res.Scan(&tBatteryHistoryData.Ts,
			&tBatteryHistoryData.Altitude,
			&tBatteryHistoryData.Latitude,
			&tBatteryHistoryData.Longitude,
			&tBatteryHistoryData.Mobile,
			&tBatteryHistoryData.InfoTime,
			&tBatteryHistoryData.Speed,
			&tBatteryHistoryData.Direction,
			&tBatteryHistoryData.BatteryId,
			&tBatteryHistoryData.TotalVoltage,
			&tBatteryHistoryData.TotalCurrent,
			&tBatteryHistoryData.Soc,
			&tBatteryHistoryData.CellQuantity,
			&tBatteryHistoryData.CellVoltageDetail,
			&tBatteryHistoryData.TempQuantity,
			&tBatteryHistoryData.TempDetailInfo,
			&tBatteryHistoryData.BMSTemp,
			&tBatteryHistoryData.ResidualCapacity,
			&tBatteryHistoryData.CurrentCapacity,
			&tBatteryHistoryData.LoopTimes,
			&tBatteryHistoryData.GsmSignalIntensity,
			&tBatteryHistoryData.LocationSatNum)
		tBatteryHistoryDataList = append(tBatteryHistoryDataList, tBatteryHistoryData)
	}
	return
}

type Student struct {
	Ts  int64 `json:"ts"`
	Sid int   `json:"sid"`
	Age int   `json:"age"`
}

// 获取从kafka过来的数值，通过json转为go对象
type ConsumerBatteryHistoryData struct {
	// AlarmDetail       string  `json:"alarmDetail"`
	// AlarmStatus       int     `json:"alarmStatus"`
	Altitude          int     `json:"altitude"`
	BMSTemp           int     `json:"bMSTemp"`
	BatteryID         string  `json:"batteryId"`
	CellQuantity      int     `json:"cellQuantity"`
	CellVoltageDetail string  `json:"cellVoltageDetail"`
	CreateTime        int64   `json:"createTime"`
	CurrentCapacity   float64 `json:"currentCapacity"`
	Direction         int     `json:"direction"`
	InfoTime          int64   `json:"infoTime"`
	Latitude          float64 `json:"latitude"`
	LifeSignal        float64 `json:"lifeSignal"`
	Longitude         float64 `json:"longitude"`
	LoopTimes         string  `json:"loopTimes"`
	Mobile            int64   `json:"mobile"`
	ResidualCapacity  float64 `json:"residualCapacity"`
	Soc               int     `json:"soc"`
	Speed             int     `json:"speed"`
	TempDetailInfo    string  `json:"tempDetailInfo"`
	TempQuantity      int     `json:"tempQuantity"`
	TotalCurrent      float64 `json:"totalCurrent"`
	TotalVoltage      float64 `json:"totalVoltage"`
}

// 多数据进行插入
func InsBatteryHistoryData(temp []string) {

	// sqlList := []ConsumerBatteryHistoryData{}
	// sqlListStr := []string{}

	// 从kafka读取出来比较全的数据
	// sqlStr := "insert into devDB.t_battery_history_data (ts, alarm_detail, alarm_status, altitude, b_m_s_temp, battery_id, cell_quantity, cell_voltage_detail, create_time, current_capacity, direction, info_time, latitude, life_signal, longitude, loop_times, mobile, residual_capacity, soc, speed, temp_detail_info, temp_quantity, total_current, total_voltage) values"
	sqlStr := "insert into devDB.t_battery_history_data (ts, altitude, b_m_s_temp, battery_id, cell_quantity, cell_voltage_detail, current_capacity, direction, info_time, latitude, life_signal, longitude, loop_times, mobile, residual_capacity, soc, speed, temp_detail_info, temp_quantity, total_current, total_voltage) values"

	// 进行拼接
	tagsql := ""

	for _, s := range temp {
		// fmt.Println(s)
		// alarmDetail := gjson.Get(s, "alarmDetail")
		// alarmStatus := gjson.Get(s, "alarmStatus")
		altitude := gjson.Get(s, "altitude")
		bMSTemp := gjson.Get(s, "bMSTemp")
		batteryId := gjson.Get(s, "batteryId")
		cellQuantity := gjson.Get(s, "cellQuantity")
		cellVoltageDetail := gjson.Get(s, "cellVoltageDetail")
		// createTime := gjson.Get(s, "createTime")
		currentCapacity := gjson.Get(s, "currentCapacity")
		direction := gjson.Get(s, "direction")
		infoTime := gjson.Get(s, "infoTime")
		latitude := gjson.Get(s, "latitude")
		lifeSignal := gjson.Get(s, "lifeSignal")
		longitude := gjson.Get(s, "longitude")
		loopTimes := gjson.Get(s, "loop_times")
		mobile := gjson.Get(s, "mobile")
		residualCapacity := gjson.Get(s, "residualCapacity")
		soc := gjson.Get(s, "soc")
		speed := gjson.Get(s, "speed")
		tempDetailInfo := gjson.Get(s, "tempDetailInfo")
		tempQuantity := gjson.Get(s, "tempQuantity")
		totalCurrent := gjson.Get(s, "totalCurrent")
		totalVoltage := gjson.Get(s, "totalVoltage")

		// fmt.Println(alarmDetail, "|", alarmStatus, "|", altitude, "|", bMSTemp, "|", batteryId, "|", cellQuantity, cellVoltageDetail, createTime, currentCapacity, direction, id, infoTime, latitude, lifeSignal, longitude, loopTimes, mobile, residualCapacity, soc, speed, tempDetailInfo, tempQuantity, totalCurrent, totalVoltage)

		tim := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
		tagsql = tagsql + "(" + tim +
			// "," + strconv.Itoa(int(alarmDetail.Num)) +
			// "," + strconv.Itoa(int(alarmStatus.Num)) +
			"," + strconv.Itoa(int(altitude.Num)) +
			"," + strconv.Itoa(int(bMSTemp.Num)) +
			"," + "\"" + batteryId.Str + "\"" +
			"," + strconv.Itoa(int(cellQuantity.Num)) +
			"," + "\"" + cellVoltageDetail.Str + "\"" +
			// "," + strconv.Itoa(int(createTime.Num)) +
			"," + strconv.FormatFloat(currentCapacity.Float(), 'f', 1, 64) +
			"," + strconv.Itoa(int(direction.Num)) +
			"," + strconv.Itoa(int(infoTime.Num)) +
			"," + strconv.FormatFloat(latitude.Float(), 'f', 6, 64) +
			"," + strconv.FormatFloat(lifeSignal.Float(), 'f', 1, 64) +
			"," + strconv.FormatFloat(longitude.Float(), 'f', 6, 64) +
			"," + loopTimes.Str +
			"," + strconv.Itoa(int(mobile.Num)) +
			"," + strconv.FormatFloat(residualCapacity.Float(), 'f', 1, 64) +
			"," + strconv.Itoa(int(soc.Num)) +
			"," + strconv.Itoa(int(speed.Num)) +
			"," + "'" + tempDetailInfo.Str + "'" +
			"," + strconv.Itoa(int(tempQuantity.Num)) +
			"," + strconv.FormatFloat(totalCurrent.Float(), 'f', 1, 64) +
			"," + strconv.FormatFloat(totalVoltage.Float(), 'f', 1, 64) + ")"

	}
	tt := sqlStr + tagsql
	// fmt.Println(tt)
	// 执行插入
	_, err := taos.SqlDB.Exec(tt)
	if err != nil {
		fmt.Printf("INSERT failed: %v", err)
	}
	// count, err := res.RowsAffected()
	// fmt.Println(count)

}

// 多数据进行插入
func InsBatteryHistoryDataStudent(temp []string) {
	start := time.Now()

	sqlStr := "insert into devDB.student (ts, sid, age) values"

	tagsql := ""
	for _, s := range temp {

		speed := gjson.Get(s, "speed")
		soc := gjson.Get(s, "soc")
		tim := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
		fmt.Println(tim)

		tagsql = tagsql + "(" + tim +
			"," + strconv.Itoa(int(speed.Num)) +
			"," + strconv.Itoa(int(soc.Num)) + ")"

	}
	tt := sqlStr + tagsql
	// fmt.Print(tt)

	res, err := taos.SqlDB.Exec(tt)
	if err != nil {
		fmt.Printf("INSERT failed: %v", err)
		// tx.Rollback()
	}
	count, err := res.RowsAffected()
	fmt.Println(count)

	// 	obj := TBatteryHistoryData{}
	// 	err := json.Unmarshal([]byte(temp), &obj)
	// 	if err != nil {
	// 		fmt.Println("json error")
	// 	}
	// 	fmt.Println(obj)

	end := time.Now()
	fmt.Println("tx.Exec 不释放连接 insert total time:", end.Sub(start).Seconds())
}

/**
 * 想实现多协程
 * 多数据并发
 * 支持事务
 */

func InsBatteryHistoryDataDemo() {
	// 执行时间
	start := time.Now()
	fmt.Println("开始时间 %n", start)
	// tx, err := taos.SqlDB.Begin()
	// if err != nil {
	// 	fmt.Printf("fail, err:%v\n", err)
	// 	return
	// }
	s := []Student{{1640262620000, 3, 500}, {1640262651695, 56, 200}, {1640262656697, 65, 100}}

	// pointList := [][]interface{}{{2, "孙悟空", 500}, {3, "猪八戒", 200}, {4, "沙悟净", 100}}
	// sqlStr := "insert into devDB.t_battery_history_data (ts, altitude, latitude, longitude, mobile, info_time, speed, direction, battery_id, total_voltage, total_current, soc, cell_quantity, cell_voltage_detail, temp_quantity, temp_detail_info, b_m_s_temp, residual_capacity, current_capacity, loop_times, gsm_signal_intensity, location_sat_num) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	sqlStr := "insert into devDB.student (ts, sid, age) values"
	// stmt, err := taos.SqlDB.Prepare(sqlStr)
	// if err != nil {
	// 	fmt.Printf("INSERT failed: %v", err)
	// 	// tx.Rollback()
	// }
	sqlList := []string{}
	sql := ""
	for i := 0; i < len(s); i++ {
		if i%5 == 0 {
			if sql != "" {
				//把上次拼接的SQL结果存储起来
				sqlList = append(sqlList, sql)
			}
			//重置SQL
			sql = sqlStr
		}
		if sql != sqlStr {
			sql = sql + ","
		}
		sql = sql + fmt.Sprintf("(%d,'%d','%d')", s[i].Ts, s[i].Sid, s[i].Age)
	}
	//把最后一次生成的SQL存储起来
	sqlList = append(sqlList, sql)
	fmt.Println(sqlList[0])

	res, err := taos.SqlDB.Exec(sqlList[0])
	if err != nil {
		fmt.Printf("INSERT failed: %v", err)
		// tx.Rollback()
	}
	// count, err := res.RowsAffected()
	// if err != nil {
	// 	fmt.Printf("INSERT failed: %v", err)
	// 	// tx.Rollback()
	// }
	fmt.Println(res)
	// defer taos.Close()
	// for _, val := range values {
	// 	_, err := stmt.Exec(val...)
	// 	if err != nil {
	// 		fmt.Printf("INSERT failed: %v", err)
	// 		// tx.Rollback()
	// 	}
	// }
	//最后释放tx内部的连接
	// tx.Commit()

	end := time.Now()
	fmt.Println("tx.Exec 不释放连接 insert total time:", end.Sub(start).Seconds())
}

func Hello() int {
	limit := 3
	offset := 0
	sqlStr := "select ts,soc from devDB.t_battery_history_data where soc = 0 " + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	fmt.Println(sqlStr)
	res, err := taos.SqlDB.Query(sqlStr)
	taos.CheckErr(err, sqlStr)
	if err != nil {
		fmt.Println("cuowu")
	}
	defer res.Close()
	var tsoc int
	for res.Next() {
		var (
			ts  time.Time
			soc int
		)
		err = res.Scan(&ts, &soc)
		taos.CheckErr(err, sqlStr)
		fmt.Println("nanosecond is correct: ", ts, soc)
		tsoc = soc
	}
	return tsoc
}

/*
 * Demo测试用例，不使用全局加载注入
 */

// func Hello2() {
// 	limit := 3
// 	offset := 0
// 	sqlStr := "select ts,soc from devDB.t_battery_history_data where soc = 0 " + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
// 	fmt.Println(sqlStr)
// 	db := taos.Sel(sqlStr)
// 	res, err := db.Query(sqlStr)
// 	taos.CheckErr(err, sqlStr)
// 	defer res.Close()
// 	for res.Next() {
// 		var (
// 			ts  time.Time
// 			soc int
// 		)
// 		err = res.Scan(&ts, &soc)
// 		taos.CheckErr(err, sqlStr)
// 		fmt.Println("nanosecond is correct: ", ts, soc)
// 		if ts.Nanosecond()%1000 > 0 {
// 			fmt.Println("nanosecond is correct: ", ts, soc)
// 		}
// 	}
// }
