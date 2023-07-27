package do

// SystemUserGroup 子系统用户信息，描述用户所属系统内的组信息，例如gm系统是系统管理员
type SystemUserGroup struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	Group   string `json:"group"` // 子系统自定义的组，用于权限管理
}

type GroupType = string

const (
	GroupType_SuperAdmin = "superadmin" // 超级管理员，系统只有一个
	GroupType_UserAdmin  = "useradmin"  // 用户管理员，可以创建用户，分配子系统
	GroupType_Visitor    = "common"     // 普通用户
)

type User struct {
	Systems   []*SystemUserGroup `json:"systems"` // admin|system_admin|...
	UserName  string             `json:"username"`
	Password  string             `json:"password"`
	Salt      string             `json:"salt"`
	CreatedAt string             `json:"created_at"`
}

func (u *User) Copy() *User {
	newUser := &User{
		Systems:   make([]*SystemUserGroup, len(u.Systems)),
		UserName:  u.UserName,
		Password:  u.Password,
		Salt:      u.Salt,
		CreatedAt: u.CreatedAt,
	}
	copy(newUser.Systems, u.Systems)
	return newUser
}

type Permission struct {
	Project string `json:"project"`
}
