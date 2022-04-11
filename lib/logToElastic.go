package lib

// import (
// 	"eform-gateway/lib"
// 	"fmt"
// 	"log"
// 	"runtime"
// 	"strings"
// )

// var (
// 	WarningLogger *log.Logger
// 	InfoLogger    *log.Logger
// 	ErrorLogger   *log.Logger
// 	Elastic       lib.Elasticsearch2
// )

// func WhereAmI(depthList ...int) (string, string, string) {
// 	var depth int
// 	if depthList == nil {
// 		depth = 1
// 	} else {
// 		depth = depthList[0]
// 	}

// 	function, file, line, _ := runtime.Caller(depth)
// 	fl := chopPath(file)
// 	fn := runtime.FuncForPC(function).Name()
// 	ln := fmt.Sprintf("%d", line)

// 	return fl, fn, ln
// }

// func chopPath(original string) string {
// 	i := strings.LastIndex(original, "/")
// 	if i == -1 {
// 		return original
// 	} else {
// 		return original[i+1:]
// 	}
// }

// func CreateLogErrorToDB(db *Elastic.Client, filename, function, line, messageCustom, messageSystem string) {
// 	// var logError models.LogError

// 	// t := time.Now()
// 	// y := fmt.Sprintf("%v", t.Year())
// 	// mstring := fmt.Sprintf("%v", t.Month())
// 	// mnumber := fmt.Sprintf("%02d", int(t.Month()))
// 	// my := mnumber + y
// 	// tabLog := os.Getenv("TABLOG_ERROR") + "_" + my

// 	// Create log error here

// }

// func CreateLogActivityToDB(db *Elastic.Client, uri string, agent string, ipaddress string, reqBody string, respBody string) {
// 	// var LogAccess models.LogAccess
// 	// t := time.Now()
// 	// y := fmt.Sprintf("%v", t.Year())
// 	// mnumber := fmt.Sprintf("%02d", int(t.Month()))
// 	// my := mnumber + y
// 	// tabActivity := os.Getenv("TABLOG_ACTIVITY") + "_" + my

// 	// create log activity

// }
