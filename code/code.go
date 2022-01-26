package code

const (
	CODE_SUCCESS      = iota + 10000 // 成功
	CODE_PARAM_ERROR                 // 参数错误
	CODE_SERVER_ERROR                // 服务错误
)

const (
	MESSAGE_SUCCESS      = "成功"
	MESSAGE_PARAM_ERROR  = "参数错误"
	MESSAGE_SERVER_ERROR = "服务器内部错误"
)
