package service

import "agricultural_product_gin/dto"

// 辅助函数
func successResult(msg string, data interface{}) *dto.Result {
	return &dto.Result{Code: 200, Msg: msg, Data: data}
}

func errorResult(code int, msg string) *dto.Result {
	return &dto.Result{Code: code, Msg: msg, Data: nil}
}
