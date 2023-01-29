package utils

import "bytes"

// StringBytesBufferJoin : 拼接字符串 bytes.Buffer模式
func StringBytesBufferJoin(con ...string) string {
	stringBytesBuffer := bytes.Buffer{}
	for _, s := range con {
		stringBytesBuffer.WriteString(s)
	}
	return stringBytesBuffer.String()
}
