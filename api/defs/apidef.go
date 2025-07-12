package defs

// requests 先将他的数据模型先声明出来
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// data model
type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     string `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}
