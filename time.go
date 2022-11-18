package util

import (
	"time"
)

func Milliseconds() int64 {
	return time.Now().UnixNano() / 1e6
}
