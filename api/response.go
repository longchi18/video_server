package main

import (
	"encoding/json"
	"github.com/longchi18/video-server/api/defs"
	"io"
	"net/http"
)

// 发送错误响应
func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

// 发送正常响应
func sendNornalResponse(w http.ResponseWriter, resp string, sc int) {
	// 把Header的头信息写入到响应中
	w.WriteHeader(sc)
	// 把响应的内容写入到客户端
	io.WriteString(w, resp)
}
