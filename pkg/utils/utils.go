package utils

import (
	"bytes"
	"time"
)

/*StringBytesBufferJoin
 * @description: 拼接字符串 bytes.Buffer模式
 * @param {...string} con
 * @return {*}
 */
func StringBytesBufferJoin(con ...string) string {
	stringBytesBuffer := bytes.Buffer{}
	for _, s := range con {
		stringBytesBuffer.WriteString(s)
	}
	return stringBytesBuffer.String()
}

/*GetDayTime
 * @description: 返回当前的时间
 * @return {*}
 */
func GetDayTime() string {
	//template := "2006:01:02"
	template := "2006-01-02 15:04:05" // 标准模板
	return time.Now().Format(template)
}

/*GetDay
 * @description: 获取当前年月日
 * @return {*}
 */
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

/*GetUnixNano
 * @description: 获取当前时间戳纳秒
 * @return {*}
 */
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

/*GetFileDay
 * @description: 获取当前年/月/日
 * @return {*}
 */
func GetFileDay() string {
	template := "2006/01/02"
	return time.Now().Format(template)
}
