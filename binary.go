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

// TODO 其实可以直接用binary.Write/Read(buf, order, &[]interface{struct{}, struct{}})方式来读取
// func FromBinary(b []byte, vs ...interface{}) {
// 	buf := bytes.NewBuffer(b)
// 	for _, v := range vs
// }
