package errcode

var (
	Success = NewError(0, "成功")
	ServiceError = NewError(10000, "服务器内部错误")
)
