package controllers

import (
	"joytool/apps/gmtool/service/dao"
)

type Controllers struct {
	Db *dao.Dao
}

func NewControllers(db *dao.Dao) *Controllers {
	return &Controllers{Db: db}
}

type CommandListReq struct {
	Name string `json:"command_server_name"`
}
