package dao

import (
	"joytool/apps/user/model/do"
	"joytool/lib/filedb"
	"time"
)

type Dao struct {
	userList *filedb.FileDB[do.User]
}

func NewDao() *Dao {
	ul, err := filedb.NewFileDb[do.User]("user", "user_list")
	if err != nil {
		panic(err)
	}

	d := new(Dao)
	d.userList = ul

	d.tryAddAdminUser()

	return d
}

var (
	adminUserName   = "admin"
	adminuserPasswd = "admin"
)

func (d *Dao) tryAddAdminUser() {
	doUser := &do.User{
		Systems: []*do.SystemUserGroup{
			{"user", true, ""},
		},
		UserName:  adminUserName,
		Password:  adminuserPasswd,
		Salt:      "sdfs",
		CreatedAt: time.Now().Format(time.DateTime),
	}
	d.userList.StoreOrReplace(doUser, doUser.UserName)
}
