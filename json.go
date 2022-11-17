package util

import (
	"encoding/json"
	"log"
)

func ToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("ToJson() error: %v", err)
		return ""
	}
	return string(b)
}

func FromJson(s string, v interface{}) {
	err := json.Unmarshal([]byte(s), v)
	if err != nil {
		log.Printf("FromJson() error: %v", err)
	}
}
