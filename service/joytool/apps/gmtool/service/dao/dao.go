package dao

import (
	"joytool/apps/gmtool/model/do"
)

type Dao struct {
	cmdServerList *fileDb[do.CommandServer]
	envList       *fileDb[do.Env]
	likeList      *fileDb[do.LikeExecCommand]
}

func NewDao() *Dao {
	csl, err := newFileDb[do.CommandServer]("command_server_list")
	if err != nil {
		panic(err)
	}
	el, err := newFileDb[do.Env]("env_list")
	if err != nil {
		panic(err)
	}
	ll, err := newFileDb[do.LikeExecCommand]("like_list")
	if err != nil {
		panic(err)
	}

	d := new(Dao)

	d.cmdServerList = csl
	d.envList = el
	d.likeList = ll

	return d
}
