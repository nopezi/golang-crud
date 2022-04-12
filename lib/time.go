package lib

import (
	"fmt"
	"time"
)

func GetTimeNow(param string) string {
	currentTime := time.Now()
	time.LoadLocation("Asia/Jakarta")

	switch param {
	case "timestime":
		return currentTime.Format("2006-01-02 15:04:05")
	case "date":
		return currentTime.Format("2006-01-02")
	case "year":
		return string(currentTime.Year())
	case "month":
		return currentTime.Month().String()
	case "day":
		return string(currentTime.Day())
	case "hour":
		return string(currentTime.Hour())
	case "minutes":
		return string(currentTime.Minute())
	case "second":
		return string(currentTime.Second())
	case "unixmicro":
		return string(currentTime.UnixMicro())
	default:
		fmt.Println("masukan parameter")
		return ""
	}

}
