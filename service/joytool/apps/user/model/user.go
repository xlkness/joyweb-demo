package model

type User struct {
	UserID      string
	Generation  int // 代，系统第一个管理员是第一代，管理员创建的为第二代等等
	UserName    string
	UserSaltPwd string
	Salt        string
	Permissions map[string]*Permission
}
