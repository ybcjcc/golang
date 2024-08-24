package xcode

import "net/http"

const (
	MinReservedCode = -1000 // 系统保留错误码最低值
	MaxReservedCode = 10000 // 系统保留错误码最大值
	ForbiddenCode   = 0     // 禁止的错误码
)

// 错误码范围定义：
//  - 10000 内为系统保留
//  - 其他业务参见：https://wiki.medlinker.com/pages/viewpage.action?pageId=17602122
var (
	// old： biz code: https://git.medlinker.com/service/common/blob/master/ecode/bizcode.go
	OK                   = &Code{httpStatus: http.StatusOK, code: 0, message: "请求成功"}                           // Deprecated
	CommonErr            = &Code{httpStatus: http.StatusBadRequest, code: 3, message: "通用错误(当你不知道该如何归类时使用此错误)"} // Deprecated
	OperDbErr            = &Code{httpStatus: http.StatusInternalServerError, code: 2, message: "操作db错误"}        // Deprecated
	ParamInvalid         = &Code{httpStatus: http.StatusBadRequest, code: 1, message: "参数错误"}                   // Deprecated
	UserNotLogin         = &Code{httpStatus: http.StatusUnauthorized, code: 101, message: "用户未登录"}              // Deprecated
	UserSessionExpired   = &Code{httpStatus: http.StatusUnauthorized, code: 102, message: "用户SESSION已过期"}       // Deprecated
	UserNotAuthorized    = &Code{httpStatus: http.StatusUnauthorized, code: 103, message: "鉴权失败，当前用户没有操作权限"}    // Deprecated
	NeedWechatAuthorized = &Code{httpStatus: http.StatusUnauthorized, code: 104, message: "需要额外的微信授权"}          // Deprecated
	UserHasNoPhone       = &Code{httpStatus: http.StatusBadRequest, code: 105, message: "当前用户未绑定手机号"}           // Deprecated
	SrvParamInvalid      = &Code{httpStatus: http.StatusInternalServerError, code: 1001, message: "服务调用参数出错"}   // Deprecated
	SrvOperDbErr         = &Code{httpStatus: http.StatusInternalServerError, code: 1002, message: "服务操作数据库出错"}  // Deprecated
	SrvCommonErr         = &Code{httpStatus: http.StatusInternalServerError, code: 1003, message: "服务通用错误"}     // Deprecated

	// old: sys code: https://git.medlinker.com/service/common/blob/master/ecode/syscode.go
	NotModified        = &Code{httpStatus: http.StatusNotFound, code: -304, message: "木有改动"}                 // Deprecated
	TemporaryRedirect  = &Code{httpStatus: http.StatusInternalServerError, code: -307, message: "撞车跳转"}      // Deprecated
	RequestErr         = &Code{httpStatus: http.StatusBadRequest, code: -400, message: "请求错误"}               // Deprecated
	OldUnauthorized    = &Code{httpStatus: http.StatusUnauthorized, code: -401, message: "未认证"}              // Deprecated
	AccessDenied       = &Code{httpStatus: http.StatusForbidden, code: -403, message: "访问权限不足"}              // Deprecated
	NothingFound       = &Code{httpStatus: http.StatusNotFound, code: -404, message: "啥都木有"}                 // Deprecated
	MethodNotAllowed   = &Code{httpStatus: http.StatusMethodNotAllowed, code: -405, message: "不支持该方法"}       // Deprecated
	Conflict           = &Code{httpStatus: http.StatusConflict, code: -409, message: "冲突"}                   // Deprecated
	Canceled           = &Code{httpStatus: http.StatusInternalServerError, code: -498, message: "客户端取消请求"}   // Deprecated
	ServerErr          = &Code{httpStatus: http.StatusInternalServerError, code: -500, message: "服务器错误"}     // Deprecated
	ServiceUnavailable = &Code{httpStatus: http.StatusTooManyRequests, code: -503, message: "过载保护,服务暂不可用"}   // Deprecated
	Deadline           = &Code{httpStatus: http.StatusTooManyRequests, code: -504, message: "服务调用超时"}        // Deprecated
	LimitExceed        = &Code{httpStatus: http.StatusTooManyRequests, code: -509, message: "超出限制"}          // Deprecated
	DbTransactionErr   = &Code{httpStatus: http.StatusInternalServerError, code: -510, message: "数据库事务发生崩溃"} // Deprecated

	// 特殊错误
	ErrorCodeFailed    = &Code{httpStatus: http.StatusInternalServerError, code: 100, message: "错误码定义错误"}
	Unauthorized       = &Code{httpStatus: http.StatusUnauthorized, code: 101, message: "请求未授权"}
	WechatUnauthorized = &Code{httpStatus: http.StatusUnauthorized, code: 104, message: "微信未授权"}
	NotBindPhone       = &Code{httpStatus: http.StatusBadRequest, code: 105, message: "未绑定手机号"}
	NoLiveVerify       = &Code{httpStatus: http.StatusBadRequest, code: 106, message: "未进行活体检测"}

	// 常规错误，前缀 1
	RequestParamError   = &Code{httpStatus: http.StatusBadRequest, code: 1001, message: "请求参数错误"}
	InternalServerError = &Code{httpStatus: http.StatusInternalServerError, code: 1003, message: "内部服务错误"}

	// DB 错误，前缀 2
	DBFailed           = &Code{httpStatus: http.StatusInternalServerError, code: 2001, message: "数据库操作失败"}
	DBTransactionError = &Code{httpStatus: http.StatusInternalServerError, code: 2002, message: "数据库事务错误"}
	DBRecordNotFound   = &Code{httpStatus: http.StatusNotFound, code: 2003, message: "无相关记录"}

	// GRPC 错误，前缀 3
	GRPCFailed         = &Code{httpStatus: http.StatusInternalServerError, code: 3001, message: "GRPC 服务错误"}
	GRPCMethodNotFound = &Code{httpStatus: http.StatusInternalServerError, code: 3002, message: "不支持的 GRPC 方法"}
	GRPCUnauthorized   = &Code{httpStatus: http.StatusInternalServerError, code: 3003, message: "GRPC 访问未授权"}
	GRPCAccessDenied   = &Code{httpStatus: http.StatusInternalServerError, code: 3004, message: "GRPC 访问权限不足"}
	GRPCCanceled       = &Code{httpStatus: http.StatusInternalServerError, code: 3005, message: "GRPC 客户端取消请求"}
	GRPCTimeout        = &Code{httpStatus: http.StatusInternalServerError, code: 3006, message: "GRPC 超时"}
)

var allCode = []*Code{
	OK,
	CommonErr,
	OperDbErr,
	ParamInvalid,
	UserNotLogin,
	UserSessionExpired,
	UserNotAuthorized,
	NeedWechatAuthorized,
	UserHasNoPhone,
	SrvParamInvalid,
	SrvOperDbErr,
	SrvCommonErr,
	NotModified,
	TemporaryRedirect,
	RequestErr,
	OldUnauthorized,
	AccessDenied,
	NothingFound,
	MethodNotAllowed,
	Conflict,
	Canceled,
	ServerErr,
	ServiceUnavailable,
	Deadline,
	LimitExceed,
	DbTransactionErr,

	ErrorCodeFailed,
	Unauthorized,
	WechatUnauthorized,
	NotBindPhone,
	NoLiveVerify,
	RequestParamError,
	InternalServerError,

	DBFailed,
	DBTransactionError,
	DBRecordNotFound,

	GRPCFailed,
	GRPCMethodNotFound,
	GRPCUnauthorized,
	GRPCAccessDenied,
	GRPCCanceled,
	GRPCTimeout,
}
