package defs

// requests 先将他的数据模型先声明出来
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// response data model
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// data model
type VideoInfo struct {
	Id           string
	AuthorId     string
	Name         string
	DisplayCtime string
}

// data model
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

// data model
type SimpleSession struct {
	Username string
	TTL      int64
}
