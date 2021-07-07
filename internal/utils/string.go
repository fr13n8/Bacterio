package utils

import (
	"strconv"
	"strings"
)

func SplitAfterIndex(str string, index byte) string {
	return str[strings.IndexByte(str, index)+1:]
}

func Contains(v []string, str string) bool {
	var has bool
	for _, param := range v {
		if strings.Contains(param, str) {
			has = true
			break
		}
	}
	return has
}

func Find(v []string, str string) string {
	for _, param := range v {
		if strings.Contains(param, str) {
			return param
		}
	}
	return ""
}

func StringToInt(v string) (int, error) {
	return strconv.Atoi(v)
}

func CmdJoin(str []string) string {
	var x string
	x, str = str[0], str[1:]
	x, str = str[len(str)-1], str[:len(str)-1]
	str = append(str, x)
	return strings.Join(str[:], " ")
}
