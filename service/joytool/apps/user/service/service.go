package service

import (
	dao2 "joytool/apps/user/service/dao"
)

type Service struct {
	dao *dao2.Dao
}

func NewService() *Service {
	dao := dao2.NewDao()
	svc := new(Service)
	svc.dao = dao
	return svc
}
