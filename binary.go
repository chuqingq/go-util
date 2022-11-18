package util

import (
	"bytes"
	"encoding/binary"
)

func ToBinary(vs ...interface{}) []byte {
	var b bytes.Buffer
	for _, v := range vs {
		binary.Write(&b, binary.LittleEndian, v)
	}
	return b.Bytes()
}

// func FromBinary(b []byte, vs ...interface{}) {
// 	buf := bytes.NewBuffer(b)
// 	for _, v := range vs
// }
