package util

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Sha256(data []byte) string {
	sha := sha256.New()
	sha.Write(data)
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func Now() int64 {
	return time.Now().Local().UnixNano() / 1e6
}

func ParseRegex(str string) string {
	var data = regexp.QuoteMeta(str)
	var array = strings.Split(data, " ")
	var result = ".*"
	for _, item := range array {
		result += item + ".*"
	}
	return result
}
