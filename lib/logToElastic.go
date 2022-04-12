package lib

import (
	"bytes"
	"context"
	"eform-gateway/models"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

type logActivity struct {
	Uri       string `json:"uri"`
	Agent     string `json:"agent"`
	Ipaddress string `json:"ipaddress"`
	Request   string `json:"request"`
	Response  string `json:"response"`
}

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
	model := models.LogAccess{}
	logs := logActivity{
		Uri:       uri,
		Agent:     agent,
		Ipaddress: ipaddress,
		Request:   reqBody,
		Response:  respBody,
	}
	body, err := json.Marshal(logs)
	if err != nil {
		fmt.Println(err.Error())
	}

	month := strings.ToLower(GetTimeNow("month"))
	req := esapi.CreateRequest{
		Index:      model.IndexLogAccess(month),
		DocumentID: UUID(false),
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := req.Do(ctx, db)
	if err != nil {
		fmt.Println("LogToElastic =>", err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println("LogToElastic=> ", res.String())
	}
	// create log activity
	// fmt.Println("==>", db)
	// fmt.Println("uri==>", uri)
	// fmt.Println("agent==>", agent)
	// fmt.Println("ipaddress==>", ipaddress)
	// fmt.Println("request==>", reqBody)
	// fmt.Println("response==>", respBody)
}
