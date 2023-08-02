package controllers

import (
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
	"joytool/apps/user/api_user"
)

func (ctl *Controllers) PermissionGroupList(ctx *model.MyContext, req *request.PermissionGroupList) {
	userInfo, err := ctl.UserSvc.GetUserInfo(nil, &api_user.GetUserInfoReq{UserName: ctx.GetUserName(), System: req.Project + "-gmtool"})
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	permissionGroupList := ctl.Db.PermissionGroupList(req.Project)

	permissionList, _ := ctl.Db.GetPermissionGroup(req.Project, userInfo.Group)

	ctx.RespSuccessJson(map[string]interface{}{
		"permission_list":       permissionList,
		"permission_group_list": permissionGroupList,
		"is_admin":              userInfo.IsAdmin,
	})
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
