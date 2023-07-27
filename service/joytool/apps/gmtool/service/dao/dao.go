package dao

import (
	"joytool/apps/gmtool/model/do"
	"joytool/lib/filedb"
)

type Dao struct {
	cmdServerList       *filedb.FileDB[do.CommandServer]
	envList             *filedb.FileDB[do.Env]
	likeList            *filedb.FileDB[do.LikeExecCommand]
	permissionGroupList *filedb.FileDB[do.PermissionGroupData]
}

func NewDao() *Dao {
	csl, err := filedb.NewFileDb[do.CommandServer]("gmtool", "command_server_list")
	if err != nil {
		panic(err)
	}
	el, err := filedb.NewFileDb[do.Env]("gmtool", "env_list")
	if err != nil {
		panic(err)
	}
	ll, err := filedb.NewFileDb[do.LikeExecCommand]("gmtool", "like_list")
	if err != nil {
		panic(err)
	}

	gl, err := filedb.NewFileDb[do.PermissionGroupData]("gmtool", "group_list")
	if err != nil {
		panic(err)
	}

	d := new(Dao)

	d.cmdServerList = csl
	d.envList = el
	d.likeList = ll
	d.permissionGroupList = gl

	return d
}
