package controllers

import (
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
	"joytool/apps/user/api_user"
)

func (ctl *Controllers) PermissionGroupList(ctx *model.MyContext, req *request.PermissionGroupList) {
	_, err := ctl.UserSvc.GetUserInfo(nil, &api_user.GetUserInfoReq{UserName: ctx.GetUserName(), System: "gmtool"})
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	list := ctl.Db.PermissionGroupList(req.Project)
	ctx.RespSuccessJson(list)
}

func (ctl *Controllers) CreatePermissionGroup(ctx *model.MyContext, req *request.PermissionGroupData) {
	data, err := ctl.Db.CreatePermissionGroup(req)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	ctx.RespSuccessJson(data)
	return
}

func (ctl *Controllers) EditPermissionGroup(ctx *model.MyContext, req *request.PermissionGroupData) {
	data, err := ctl.Db.EditPermissionGroup(req)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	ctx.RespSuccessJson(data)
	return
}

func (ctl *Controllers) DeletePermissionGroup(ctx *model.MyContext, req *request.DeletePermissionGroup) {
	data, err := ctl.Db.DeletePermissionGroup(req)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	ctx.RespSuccessJson(data)
	return
}
