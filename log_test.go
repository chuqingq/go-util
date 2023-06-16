package util

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	for {
		Logger().Debugf("this is debug log")
		Logger().Infof("this is info log")
		Logger().Warnf("this is warn log")
		Logger().Errorf("this is error log")
		time.Sleep(time.Second)
	}
}
