package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/longchi18/video-server/api/dbops"
	"github.com/longchi18/video-server/api/defs"
	"github.com/longchi18/video-server/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// io.WriteString(w, "Create User Handler")
	// 读取请求体中的数据
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNornalResponse(w, string(resp), 201)
		// sendNornalResponse(w, string(resp), http.StatusCreated)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
