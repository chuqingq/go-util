package util

import (
	"time"
)

func Milliseconds() int64 {
	return time.Now().UnixMilli()
	// return time.Now().UnixNano() / 1e6
}

const HandyLayout = "2006-01-02 15:04:05.000"

func TimeNowString() string {
	return time.Now().Format(HandyLayout)
}
