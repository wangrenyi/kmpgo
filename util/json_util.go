package util

import (
	"encoding/json"
	"errors"
)

//common message req
func Unmarshal(data string, model interface{}) error {
	if len(data) == 0 {
		return nil
	}
	//判断是否是指针类型
	if err := json.Unmarshal([]byte(data), model); err != nil {
		return errors.New("json fail to Unmarshal,data: " + data)
	}
	return nil
}

//common message resp
func Marshal(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("data fail to Marshal,data: " + "")
	}
	return bytes, nil
}
