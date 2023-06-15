package util

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapture(t *testing.T) {
	a := "a,b,c"
	r := Capture(a, ",", ",")
	assert.Equal(t, r, "b")
}

func TestCapture2(t *testing.T) {
	a := "a,b,c"
	r := Capture(a, "^", "c")
	log.Printf("r: %v", r)
	assert.Equal(t, r, "a,b,")

	r = Capture(a, "b", "$")
	assert.Equal(t, r, ",c")
}

func TestCapture3(t *testing.T) {
	a := "total 0\nlrwxrwxrwx 1 root root 13 Jul 21 10:45 mmc-G8GTF4R_0x74dbd49e -> ./../mmcblk2\n"
	r := Capture(a, "0x", " ")
	assert.Equal(t, r, "74dbd49e")

	log.Printf("r: %v", r)
	h, err := hex.DecodeString(r)
	log.Printf("h: %v, %v", h, err)
	assert.NoError(t, err)
}
