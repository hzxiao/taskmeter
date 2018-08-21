package util

import (
	"crypto/sha256"
	"fmt"
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
