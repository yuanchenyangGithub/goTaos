/*
 * In this test program, we'll create a database and import 1000 records
 * with unsigned integers
 *
 * Authored by <Huo Linhe> linhe.huo@gmail.com
 */
package taos

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/taosdata/driver-go/v2/taosSql"

	confs "disutaos/config"
)

type config struct {
	hostName   string
	serverPort string
	user       string
	password   string
	dbName     string
}

var configPara config

var conf = confs.GetConfig()

var taosDriverName = "taosSql"
var url string
var SqlDB sql.DB

func init() {
	// var err error
	// 测试环境
	// flag.StringVar(&configPara.hostName, "h", "izwz9cdt", "The host to connect to TDengine server.")
	flag.StringVar(&configPara.hostName, "h", conf.TaosDb.HostName, "The host to connect to TDengine server.")
	flag.StringVar(&configPara.serverPort, "p", conf.TaosDb.ServerPort, "The TCP/IP port number to use for the connection to TDengine server.")
	flag.StringVar(&configPara.user, "u", conf.TaosDb.User, "The TDengine user name to use when connecting to the server.")
	flag.StringVar(&configPara.password, "P", conf.TaosDb.Password, "The password to use when connecting to the server.")
	flag.StringVar(&configPara.dbName, "d", conf.TaosDb.DbName, "Destination database.")
	flag.Parse()

	url = "root:taosdata@/tcp(" + configPara.hostName + ":" + configPara.serverPort + ")/"
	fmt.Println(url)

	// SqlDB, err := sql.Open(taosDriverName, url)
	// if err != nil {
	// 	fmt.Printf("Open database error: %s\n", err)
	// 	os.Exit(1)
	// }
	// err = SqlDB.Ping()
	// if err != nil {
	// 	fmt.Printf("Open database error: %s\n", err)
	// }

}

func Close() {
	SqlDB.Close()
}

func PrintAllArgs() {
	fmt.Printf("============= args parse result: =============\n")
	fmt.Printf("hostName:             %v\n", configPara.hostName)
	fmt.Printf("serverPort:           %v\n", configPara.serverPort)
	fmt.Printf("usr:                  %v\n", configPara.user)
	fmt.Printf("password:             %v\n", configPara.password)
	fmt.Printf("dbName:               %v\n", configPara.dbName)
	fmt.Printf("================================================\n")
}

func Sel(sqlStr string) sql.DB {
	Bb, err := sql.Open(taosDriverName, url)
	if err != nil {
		fmt.Printf("Open database error: %s\n", err)
		os.Exit(1)
	}
	defer Bb.Close()

	fmt.Printf("- %s\n", sqlStr)
	return *Bb
}

/*
 * 分页查询
 */
func SelPage(sqlStr string, limit int, offset int) {

}

func InitTaoSql() (err error) {
	fmt.Println("sql进来了")
	SqlDB1, err := sql.Open(taosDriverName, url)
	if err != nil {
		fmt.Printf("Open database error: %s\n", err)
		os.Exit(1)
	}
	SqlDB = *SqlDB1
	err = SqlDB1.Ping()
	if err != nil {
		fmt.Println("ping失败")
	}

	// fmt.Println(SqlDB)
	return nil
}

func CheckErr(err error, prompt string) {
	if err != nil {
		fmt.Errorf("ERROR: %s\n", prompt)
		panic(err)
	}
}
