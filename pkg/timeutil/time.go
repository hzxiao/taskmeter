package timeutil

import (
	"time"
)

func Now() int64 {
	return time.Now().Local().UnixNano() / int64(time.Millisecond)
}

func GetDateString(timestamp int64) string {
	tm := time.Unix(0, timestamp*1e6)
	return tm.Format("02/01/2006 15:04:05")
}
