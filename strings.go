package util

import (
	"strings"
)

func Capture(s, left, right string) string {
	r := strings.SplitN("^"+s, left, 2)
	if len(r) != 2 {
		return ""
	}
	// log.Printf("r: %v", r)
	s = r[1]
	r = strings.SplitN(s+"$", right, 2)
	if len(r) != 2 {
		return ""
	}
	// log.Printf("r: %v", r)
	return r[0]
}
