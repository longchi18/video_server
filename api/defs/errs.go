package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"` // 系统内部用ErrorCode，对外暴露的错误码为error_code字段值。
}

type ErrResponse struct {
	HttpSC int `json:"http_status"`
	Error  Err `json:"error"`
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
)
