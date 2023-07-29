package controller

// 任意类型的数据
type any = interface{}

// 返回给视图的结果的包装器
type ResultWrapper struct {

	/*
		2000 代表成功
		4000 代表失败
	*/
	Code int32
	Msg  string
	Data any
}

// 返回错误响应对象的快捷方法
func ServerFailed(msg string) *ResultWrapper {
	wrapper := ResultWrapper{
		Code: 4000,
		Msg:  msg,
		Data: nil,
	}
	return &wrapper
}

// 返回正确响应对象的快捷方法
// data是要返回给前端的数据
func ServerSuccess(data any) *ResultWrapper {
	wrapper := ResultWrapper{
		Code: 2000,
		Msg:  "success",
		Data: data,
	}
	return &wrapper
}
