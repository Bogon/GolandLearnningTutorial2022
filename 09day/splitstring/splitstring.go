package splitstring

import (
	"strings"
)

// 切割字符串
// example
// abc b => [a, c]
func Split(s string, sep string) []string {
	//return strings.Split(s, sep)
	var ret []string
	index := strings.Index(s, sep)
	for index >= 0 {
		ret = append(ret, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return ret
}
