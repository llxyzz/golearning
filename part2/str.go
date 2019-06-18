package part2

import (
	"bytes"
	"strconv"
	"strings"
)

func BaseName(s string) string {

	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//将一个整型的数据用逗号分隔，如123456789 --> 123,456,789
func SplitInt(n string) string {
	len := len(n)
	if len <= 3 {
		return n
	}
	return SplitInt(n[:len-3]) + "," + n[len-3:]
}

//把整型数组准换为字符串数组
func IntsToString(ints []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range ints {
		if v > 0 {
			buf.Write([]byte(strconv.Itoa(v)))
			if i < len(ints)-1 {
				buf.WriteString(",")
			}
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

//将一个整型的数据用逗号分隔，如123456789 --> 123,456,789
func SplitIntByBuffer(n string) string {
	l := len(n)
	if l <= 3 {
		return n
	}

	var buf bytes.Buffer
	a := l % 3
	buf.WriteString(n[:a])
	for i := a; i < l; i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(n[i : i+3])
	}
	return buf.String()
}
