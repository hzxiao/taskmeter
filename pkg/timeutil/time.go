package timeutil

import (
	"time"
)

func Now() int64 {
	return time.Now().Unix() / int64(time.Millisecond)
}
