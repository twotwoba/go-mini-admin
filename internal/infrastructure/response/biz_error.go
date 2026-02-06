package response

// ------------------------- 错误码 -------------------------
// 1位模块 + 4位错误类型
const (
	CodeSuccess = 0
	// 通用错误
	CodeParamError    = 40000 // 请求参数错误
	CodeUnauthorized  = 40001 // 未授权
	CodeParamValidate = 40002 // 参数校验错误
	CodeForbidden     = 40003 // 禁止访问
	CodeNotFound      = 40004 // 资源不存在
	CodeServerError   = 50000 // 服务器内部错误
	// 用户模块
	CodeUserNotExist = 20001 // 用户不存在
	// 其他模块 3xxxx...
)

var errMsgMap = map[int]string{
	CodeParamError:    "请求参数错误",
	CodeParamValidate: "参数校验错误",
	CodeUnauthorized:  "未授权",
	CodeForbidden:     "禁止访问",
	CodeNotFound:      "资源不存在",
	CodeServerError:   "服务器内部错误",
	CodeUserNotExist:  "用户不存在",
}

func GetErrMsg(code int) string {
	return errMsgMap[code]
}
