package controllers

import (
	"joytool/apps/gmtool/service/dao"
	"joytool/apps/user/api_user"
)

type Controllers struct {
	Db      *dao.Dao
	UserSvc api_user.UserServiceInterface
}

func NewControllers(db *dao.Dao) *Controllers {
	userSvc := api_user.NewUserServiceInstance()
	return &Controllers{Db: db, UserSvc: userSvc}
}

type CommandListReq struct {
	Name string `json:"command_server_name"`
}
