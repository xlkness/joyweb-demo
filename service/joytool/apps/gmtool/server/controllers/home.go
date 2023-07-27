package controllers

import (
	"fmt"
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
	"joytool/apps/user/api_user"
	"sort"
)

func (ctl *Controllers) HomeView(ctx *model.MyContext, params *request.HomeView) {
	if params.Project == "" {
		ctx.RespFailMessage(300, fmt.Sprintf("project empty!please input project"))
		return
	}

	_, err := ctl.UserSvc.GetUserInfo(nil, &api_user.GetUserInfoReq{UserName: ctx.GetUserName()})
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	fmt.Printf("receive gmtool home view, user:%+v\n", ctx.GetUserName())

	commandServerList, _ := ctl.Db.QueryCmdServerList(params.Project)
	sort.SliceStable(commandServerList, func(i, j int) bool {
		return commandServerList[i].Name > commandServerList[j].Name
	})

	envs, _ := ctl.Db.QueryEnvList(params.Project)
	sort.SliceStable(envs, func(i, j int) bool {
		if envs[i].Index == envs[j].Index {
			return envs[i].Name > envs[j].Name
		}
		return envs[i].Index < envs[j].Index
	})

	likeList, _ := ctl.Db.LikeCommandList(params.Project, "test")
	sort.SliceStable(likeList, func(i, j int) bool {
		return likeList[i].Command.Date < likeList[j].Command.Date
	})

	permissionGroupList := ctl.Db.PermissionGroupList(params.Project)

	ctx.RespSuccessJson(map[string]interface{}{
		"command_server_list":   commandServerList,
		"env_list":              envs,
		"like_list":             likeList,
		"permission_group_list": permissionGroupList,
	})
}
