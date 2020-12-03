package models

import (
	"log"
	"time"
)

// 自定义模板函数
func Hello(in string) (out string) {
	out = in + "world"
	return out
}

// 把时间戳转换成时间格式
func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 把时间格式化的字符串转换成时间戳
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		log.Panic(err)
		return 0
	}
	return t.Unix()
}
