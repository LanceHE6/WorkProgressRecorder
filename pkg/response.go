package pkg

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewResponse
//
//	@Description: 创建一个响应
//	@param code 响应码
//	@param msg 响应信息
//	@param data 响应数据
//	@return *Response 响应对象
func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// SuccessResponse
//
//	@Description: 创建一个成功的响应
//	@param data 响应数据
//	@return *Response 响应对象
func SuccessResponse(data interface{}) *Response {
	return NewResponse(0, "success", data)
}

// ErrorResponse
//
//	@Description: 创建一个失败的响应
//	@param code 响应码
//	@param msg 响应信息
//	@return *Response 响应对象
func ErrorResponse(code int, msg string) *Response {
	return NewResponse(code, msg, nil)
}
