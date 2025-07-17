package dbops

import (
	"database/sql"
	"github.com/longchi18/video-server/api/defs"
	"log"
	"strconv"
	"sync"
)

// 写入session信息到数据库中
func InterSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`INSERT INTO sessions(session_id, TTL, login_name) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 从数据库中读取session信息
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare(`SELECT TTL, login_name FROM sessions WHERE session_id = ?`)
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	// var ttlint int64
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

// 从数据库中读取所有session信息 返回一个map[string]*defs.SimpleSession

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrieve sessions error: %s", err)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

// 删除session信息
func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare(`DELETE FROM sessions WHERE session_id = ?`)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err = stmtOut.Query(sid); err != nil {
		return err
	}
	// defer stmtOut.Close()
	return err
}

// func UpdateSession(sid string, ttl int64) error {
// 	ttlstr := strconv.FormatInt(ttl, 10)
// 	stmtOut, err := dbConn.Prepare(`UPDATE sessions SET TTL = ? WHERE session_id = ?`)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = stmtOut.Exec(ttlstr, sid)
// 	defer stmtOut.Close()
// 	return err
// }
