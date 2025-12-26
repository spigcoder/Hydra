package errs

// 一个错误码具有三个属性：错误编码、错误码名称、HTTP 状态码。
// 错误码不能被程序任意修改，所以应该定义为常量而非变量，但由于其还具有其它额外的属性，错误码名称和对应的 HTTP 状态码，
// 所以将其对应为一个 int 类型，代表其对应的索引，然后将其它属性定义在数组中。
//
// 错误码的定义方法参考自：https://github.com/encoredev/encore.dev/blob/v1.46.1/beta/errs/codes.go

// ErrCode 为错误编码
type ErrCode int

// 错误码枚举值，此为定义错误码对应的编号
// Note：每添加一个错误码都需要在 codeNames 和 codeStatus 中定义对应的值
const (
	// Unknown 代表服务未知的错误，对应 500 Internal Server Error
	Unknown ErrCode = 1

	// BadRequest 代表传入的参数无效，对应 400 Bad Request
	BadRequest ErrCode = 2

	// DatabaseError 代表数据库操作错误，使用 500 Internal Server Error
	DatabaseError ErrCode = 3

	// NotFound 代表为资源不存在，对应 404 Not Found
	NotFound ErrCode = 4

	// Unauthorized 代表缺乏身份验证凭证，不允许接下来的操作，对应 401 Unauthorized
	Unauthorized ErrCode = 5

	// Forbidden 代表拒绝访问，对应 403 Forbidden
	Forbidden ErrCode = 6

	// TooManyRequests 代表请求过多，对应 429 Too Many Requests
	TooManyRequests ErrCode = 7

	// Conflict 代表资源冲突，对应 409 Conflict
	Conflict ErrCode = 8

	// URINotFound 表示我们找不到注册的 URI
	URINotFound ErrCode = 9

	// RandomError 表示生成随机数错误
	RandomError ErrCode = 10

	// ParseError 表示各种序列化或反序列化失败, 或者类似 url.Parse 这样
	ParseError ErrCode = 11

	// ErrorCode 表示 请求 code 错误
	ErrorCode ErrCode = 12

	// ErrorToken 表示 token 或者 refreshToken 错误
	ErrorToken ErrCode = 13

	// ErrorInternal 表示内部错误，一般指是下游的错误
	ErrorInternal ErrCode = 14
)

// String 返回错误编码实际对应的错误码
func (c ErrCode) String() string {
	return codeNames[c]
}

// StatusCode 用于返回错误码对应的 HTTP 状态码，对于 API 响应非常友好，
// 状态码代表的是一类错误，一个状态码下可能有多个错误码
func (c ErrCode) StatusCode() int {
	return codeStatus[c]
}

// [...] 表示由编译器根据初始化值的数量来确定数组长度
// codeNames 为错误码的实际名称，用于返回给应用
var codeNames = [...]string{
	Unknown:         "unknown",
	BadRequest:      "bad_request",
	DatabaseError:   "database_error",
	NotFound:        "not_found",
	Unauthorized:    "unauthorized",
	Forbidden:       "forbidden",
	TooManyRequests: "too_many_requests",
	Conflict:        "conflict",
	URINotFound:     "bad_uri",
	RandomError:     "random_error",
	ParseError:      "parse_error",
	ErrorCode:       "error_code",
	ErrorToken:      "error_token",
	ErrorInternal:   "error_internal",
}

// codeStatus 为错误码对应的 HTTP 状态码
// 响应状态码参考：https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status
var codeStatus = [...]int{
	Unknown:         500,
	BadRequest:      400,
	DatabaseError:   500,
	NotFound:        404,
	Unauthorized:    401,
	Forbidden:       403,
	TooManyRequests: 429,
	Conflict:        409,
	URINotFound:     400,
	RandomError:     500,
	ParseError:      500,
	ErrorCode:       400,
	ErrorToken:      400,
	ErrorInternal:   500,
}
