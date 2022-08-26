package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (

	// CURD 常用业务状态码
	CurdStatusOkCode  int    = 200
	CurdStatusOkMsg   string = "Success"
	CurdCreatFailCode int    = 500
	CurdCreatFailMsg  string = "操作失败"
)
