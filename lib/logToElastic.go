package lib

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func WhereAmI(depthList ...int) (string, string, string) {

	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}

	function, file, line, _ := runtime.Caller(depth)
	fl := chopPath(file)
	fn := runtime.FuncForPC(function).Name()
	ln := fmt.Sprintf("%d", line)

	return fl, fn, ln
}

func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

func CreateLogErrorToDB(db ElasticSearch2, filename, function, line, messageCustom, messageSystem string) {
	// var logError models.LogError

	// t := time.Now()
	// y := fmt.Sprintf("%v", t.Year())
	// mstring := fmt.Sprintf("%v", t.Month())
	// mnumber := fmt.Sprintf("%02d", int(t.Month()))
	// my := mnumber + y
	// tabLog := os.Getenv("TABLOG_ERROR") + "_" + my

	// Create log error here

}

func CreateLogActivityToDB(db *elasticsearch.Client, uri string, agent string, ipaddress string, reqBody string, respBody string) {
	// var LogAccess models.LogAccess
	// t := time.Now()
	// y := fmt.Sprintf("%v", t.Year())
	// mnumber := fmt.Sprintf("%02d", int(t.Month()))
	// my := mnumber + y
	// tabActivity := os.Getenv("TABLOG_ACTIVITY") + "_" + my

	// create log activity
	fmt.Println(db)
	fmt.Println(uri)
	fmt.Println(agent)
	fmt.Println(ipaddress)
	fmt.Println(reqBody)
	fmt.Println(respBody)

}
