package request

import "joytool/apps/user/model/do"

type UserBaseInfo struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type CreateUser struct {
	BaseInfo  *UserBaseInfo         `json:"base_info"`
	GroupList []*do.SystemUserGroup `json:"group_list"`
}

type LoginData struct {
	BaseInfo *UserBaseInfo `json:"base_info"`
}

type DeleteUser struct {
	UserName string `json:"username"`
}

type ListUser struct {
	System string `json:"system"`
}
