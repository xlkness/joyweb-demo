package service

import "joytool/apps/gmtool/service/dao"

type Service struct {
	Db *dao.Dao
}

func NewService() *Service {
	d := dao.NewDao()
	svc := new(Service)
	svc.Db = d
	return svc
}
