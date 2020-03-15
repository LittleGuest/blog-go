package response

import (
	"blog/consts"
	"encoding/json"
)

// Resp 响应
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// PageInfo 分页信息
type PageInfo struct {
	Offset int64       `json:"offset"`
	Limit  int64       `json:"limit"`
	Total  int64       `json:"total"`
	Data   interface{} `json:"data"`
}

func encode(resp Resp) string {
	res, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(res)
}

// 返回成功信息
func Success(data interface{}) string {
	return encode(Resp{
		Code: consts.StatusOk,
		Data: data,
	})
}

// 返回错误信息
func Error(code int, msg string) string {
	return encode(Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// Page 返回分页信息
func Page(page PageInfo) string {
	return encode(Resp{
		Code: consts.StatusOk,
		Data: page,
	})
}
