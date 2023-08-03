package server

import (
	"fmt"
	"joytool/apps/user/api_user"
	"joytool/apps/user/model"
	doUser "joytool/apps/user/model/do"
	"joytool/apps/user/model/request"
	"joytool/apps/user/model/response"
	"joytool/lib/token"
	"log"
	"time"
)

func (s *Server) CreateUser(ctx *model.MyContext, userData *request.CreateUser) {
	doUser, err := s.svc.Dao.CreateUser(userData)
	if err != nil {
		ctx.RespSuccessMessage(err.Error())
		return
	}
	ctx.RespSuccessJson(userInfoDo2Dto(ctx, doUser, userData.GroupList[0].Name))
}

func (s *Server) EditUser(ctx *model.MyContext, userData *request.CreateUser) {
	doUser, err := s.svc.Dao.EditUser(userData)
	if err != nil {
		ctx.RespFailMessage(301, err.Error())
		return
	}
	ctx.RespSuccessJson(userInfoDo2Dto(ctx, doUser, userData.GroupList[0].Name))
}

func (s *Server) DeleteUser(ctx *model.MyContext, deleteData *request.DeleteUser) {
	_, err := s.svc.Dao.DeleteUser(deleteData)
	if err != nil {
		ctx.RespSuccessMessage(err.Error())
		return
	}
	ctx.RespSuccessMessage("ok")
}

const (
	tokenExpireDura = time.Hour * 24 * 30
	//tokenExpireDura = time.Second * 30
)

func (s *Server) Login(ctx *model.MyContext, loginData *request.LoginData) {
	log.Printf("receive login data:%v,%v", loginData.BaseInfo.UserName, loginData.BaseInfo.PassWord)

	// 校验用户名密码
	userData, find, err := s.svc.Dao.LoginVerify(loginData)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	if !find {
		ctx.RespFailMessage(300, "数据库不存在用户")
		return
	}

	tokenStr, err := token.GetToken(userData.UserName, tokenExpireDura)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	checkedToken, err := token.ValidToken(tokenStr)
	if err != nil {
		fmt.Printf("valid token(%v) not pass, error:%v\n", tokenStr, err)
	} else {
		checkedClaims := checkedToken.Claims.(*token.RegisteredTokenClaims)
		fmt.Printf("valid token(%v) pass, user:%v, expire:%v, issuer:%v\n",
			tokenStr, checkedClaims.User, checkedClaims.ExpiresAt, checkedClaims.Issuer)
	}

	userInfo := &api_user.GetUserInfoRes{}
	err = s.svc.GetUserInfo(nil, &api_user.GetUserInfoReq{UserName: userData.UserName, System: "user"}, userInfo)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	resp := map[string]interface{}{
		"username":  userData.UserName,
		"token":     tokenStr,
		"expire":    tokenExpireDura.Seconds(),
		"expire_at": time.Now().Add(tokenExpireDura).Format(time.DateTime),
		"is_admin":  userInfo.IsAdmin,
	}
	ctx.RespSuccessJson(resp)
}

func (s *Server) ListUsers(ctx *model.MyContext, params *request.ListUser) {
	list, _ := s.svc.Dao.UserListBySystem(params.System)
	username := ctx.GetUserName()
	for _, v := range list {
		if v.UserName == username {
			v.IsMyInfo = true
		}
		if v.IsAdmin {
			v.IsAdminStr = "是"
		} else {
			v.IsAdminStr = "否"
		}
	}
	ctx.RespSuccessJson(list)
}

func (s *Server) Logout(ctx *model.MyContext) {
	fmt.Printf("user logout:%v\n", ctx.GetUserName())
	ctx.RespSuccessMessage("ok")
}

func userInfoDo2Dto(ctx *model.MyContext, do *doUser.User, system string) *response.UserInSystem {
	dto := &response.UserInSystem{
		UserName:  do.UserName,
		CreatedAt: do.CreatedAt,
	}
	for _, v := range do.Systems {
		if v.Name == system {
			dto.IsAdmin = v.IsAdmin
			dto.Group = v.Group
			break
		}
	}

	username := ctx.GetUserName()
	if dto.UserName == username {
		dto.IsMyInfo = true
	}
	if dto.IsAdmin {
		dto.IsAdminStr = "是"
	} else {
		dto.IsAdminStr = "否"
	}

	return dto
}
