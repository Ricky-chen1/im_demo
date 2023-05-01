package util

import "time"

//将时间戳转化为时间字符串
func ConvUnixToTime(unix int64) string {
	goTime := "2006-01-02 15:04:05"
	return time.Unix(unix, 0).Format(goTime)
}

//返回当下时间时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}
