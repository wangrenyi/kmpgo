package model

import (
	"encoding/json"

	"github.com/wangrenyi/kmpgo/log"
)

// 公共消息码
const (
	Success     = 0
	ServerError = 500
	ParamsError = 401

	DBError            = 50002
	NoData             = 50003
	BigJsonError       = 50004
	JsonUnmarshalError = 50005
	UpdateError        = 50006
	UpdateNoRecord     = 50007
	SaveError          = 50008
	DeleteError        = 50009
	IllegalOperation   = 50010
	ForbiddenOperation = 50011
)

type StandardResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var statusText = map[int]string{
	Success:     "成功",
	ServerError: "服务器错误",
	ParamsError: "参数错误",
}

func StatusText(code int) string {
	return statusText[code]
}

func GenerateJSON(code int, data interface{}) (res string) {
	rsp := &StandardResponse{
		Code: code,
		Msg:  statusText[code],
		Data: data,
	}

	rspBytes, err := json.Marshal(rsp)
	if err != nil {
		log.Infof("GenerateJSON", "json.Marshal", "err: %v", err)
	}

	return string(rspBytes)
}
