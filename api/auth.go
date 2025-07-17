package main

import (
	_ "github.com/longchi18/video-server/api/defs"
	"github.com/longchi18/video-server/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"
var HEADER_FIELD_USERID = "X-User-Id"
var HEADER_FIELD_TOKEN = "X-Access-Token"

// 验证用户session
func validateUserSession(r *http.Request) bool {
	//获取session id
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	//验证session是否过期
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	// 设置用户信息到请求头中，便于后续使用
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

// 验证用户是否登录
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		// sendErrorResponse(w, r, "用户未登录")
		// sendErrorResponse()
		return false
	}
	return true
}
