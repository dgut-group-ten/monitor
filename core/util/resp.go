package util

import (
	"encoding/json"
	"log"
)

// JSONBytes : 对象转json格式的二进制数组
func JSONBytes(resp interface{}) []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}
