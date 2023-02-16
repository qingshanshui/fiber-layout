package utils

import (
	"bytes"
	"time"
)

// StringBytesBufferJoin : 拼接字符串 bytes.Buffer模式
func StringBytesBufferJoin(con ...string) string {
	stringBytesBuffer := bytes.Buffer{}
	for _, s := range con {
		stringBytesBuffer.WriteString(s)
	}
	return stringBytesBuffer.String()
}

// GetDayTime 返回当前的时间 :		2022-07-27 09:06:05
func GetDayTime() string {
	//template := "2006:01:02"
	template := "2006-01-02 15:04:05" // 标准模板
	return time.Now().Format(template)
}

// GetDay 获取当前年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// GetUnixNano 获取当前时间戳纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// GetFileDay 获取当前年/月/日
func GetFileDay() string {
	template := "2006/01/02"
	return time.Now().Format(template)
}
