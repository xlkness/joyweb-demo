package service

import (
	"context"
	"fmt"
	"joytool/apps/user/api_user"
	dao2 "joytool/apps/user/service/dao"
)

type Service struct {
	Dao *dao2.Dao
}

func (s *Service) GetUserInfo(ctx context.Context, req *api_user.GetUserInfoReq, res *api_user.GetUserInfoRes) error {
	userData, find := s.Dao.GetUserInfo(req.GetUserName(), req.GetSystem())
	if !find {
		return fmt.Errorf("数据库找不到用户信息：%v", req.GetUserName())
	}
	res.UserName = userData.UserName
	res.IsAdmin = userData.Systems[0].IsAdmin
	res.Group = userData.Systems[0].Group
	return nil
}

func NewService() *Service {
	dao := dao2.NewDao()
	svc := new(Service)
	api_user.NewUserHandler(svc)
	svc.Dao = dao
	return svc
}
