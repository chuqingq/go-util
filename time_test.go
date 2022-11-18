package util

import (
	"testing"
	"time"
)

func TestTimeLayout(t *testing.T) {
	println(time.Now().Format(HandyLayout))
	// 2022-11-18 09:45:43.451
}
