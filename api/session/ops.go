package session

import (
	"github.com/longchi18/video-server/api/dbops"
	"github.com/longchi18/video-server/api/defs"
	"github.com/longchi18/video-server/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

// init 初始化会话管理器，加载所有会话数据到内存中。
func init() {
	sessionMap = &sync.Map{}
}

// nowInMilli 返回当前时间的时间戳（毫秒）。
func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

// deleteExpiredSessions 删除过期的会话数据。
func deleteExpiredSessions(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// LoadSessionFromDB 从数据库加载会话数据到内存中。
func LoadSessionFromDB() {
	// 从数据库加载session数据到内存中
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

// CreateSession 创建一个新的会话，并返回会话ID和用户信息
func GenerateNewSessionId(un string) string {
	// 生成新的会话ID，例如使用UUID
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 // 设置会话过期时间，例如半时后过期
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InterSession(id, ttl, un)
	return id
}

// IsSessionExpired 检查会话是否过期，如果已过期则删除并返回true。
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSessions(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}

/*
type SessionManager interface {
	Set(key, value interface{}) error // 设置session
	Get(key interface{}) interface{}  // 获取session
	Delete(key interface{}) error     // 删除session
	SessionID() string                // 会话ID
	SessionRegenerateID() string      // 重新生成会话id
	Flush() error                    // 清除session值
	Save() error                     // 保存session
	Close() error                    // 关闭session
	Options(o *Options)              // 设置session选项
	Options() *Options               // 获取session选项
	IsNew() bool                    // 是否是新会话
	ID() string                    // 会话id
	Values() map[string]interface{} // 所有值
	SetValues(values map[string]interface{}) // 设置所有值
	Clear()                            // 清除所有值
	ClearValues()                    // 清除所有值
	ClearOptions()                   // 清除所有选项
}

*/
