package server

import (
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
	"joytool/apps/gmtool/server/controllers"
	"reflect"
)

func (s *Server) initRouter() {
	ctl := controllers.NewControllers(s.svc.Db)

	s.ctl = ctl

	s.registerGet("/gmtool", "gmtool", request.HomeView{}, ctl.HomeView, false)

	s.registerPost("/addcommandserver", "添加指令服务器", request.CommandServerData{}, ctl.AddCommandServer, true)
	s.registerPost("/editcommandserver", "编辑指令服务器", request.CommandServerData{}, ctl.EditCommandServer, true)
	s.registerPost("/deletecommandserver", "编辑指令服务器", request.DeleteCommandServer{}, ctl.DeleteCommandServer, true)
	s.registerPost("/deleteexechistory", "删除执行历史", request.DeleteExecHistory{}, ctl.DeleteExecHistory, true)

	s.registerPost("/addenv", "添加环境", request.EnvData{}, ctl.AddEnv, true)
	s.registerPost("/editenv", "编辑环境", request.EnvData{}, ctl.EditEnv, true)
	s.registerPost("/deleteenv", "编辑环境", request.DeleteEnv{}, ctl.DeleteEnv, true)

	s.registerGet("/commandlist", "get command list", request.CommandList{}, ctl.CommandList, false)
	s.registerPost("/execcommand", "exec command list", request.ExecCommand{}, ctl.ExecCommand, false)

	s.registerPost("/likecommand", "exec command list", request.ExecCommand{}, ctl.LikeCommand, true)
	s.registerPost("/dislikecommand", "exec command list", request.DeleteLikeCommand{}, ctl.DislikeCommand, true)

	//s.engine.GetWithStructParams("/listpermissiongroup", "exec command list", request.PermissionGroupList{}, ctl.PermissionGroupList)
	s.registerGet("/listpermissiongroup", "exec command list", request.PermissionGroupList{}, ctl.PermissionGroupList, false)
	s.registerPost("/createpermissiongroup", "exec command list", request.PermissionGroupData{}, ctl.CreatePermissionGroup, true)
	s.registerPost("/editpermissiongroup", "exec command list", request.PermissionGroupData{}, ctl.EditPermissionGroup, true)
	s.registerPost("/deletepermissiongroup", "exec command list", request.DeletePermissionGroup{}, ctl.DeletePermissionGroup, true)
}

func (s *Server) registerGet(path, desc string, dataReceiver interface{}, handler interface{}, needAdmin bool) {
	s.registerMethod("get", path, desc, dataReceiver, handler, needAdmin)
}

func (s *Server) registerPost(path, desc string, dataReceiver interface{}, handler interface{}, needAdmin bool) {
	s.registerMethod("post", path, desc, dataReceiver, handler, needAdmin)
}

func (s *Server) registerMethod(method, path, desc string, dataReceiver interface{}, handler interface{}, needAdmin bool) {
	newHandler := handler
	if needAdmin {
		newHandler = func(ctx *model.MyContext, params any) {
			project := reflect.ValueOf(params).Elem().FieldByName("Project").String()
			if ok, err := s.ctl.AdminPermission(project, ctx); err != nil {
				ctx.RespFailMessage(300, err.Error())
				return
			} else if !ok {
				ctx.RespFailMessage(300, "需要管理员权限")
				return
			}
			reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(params)})
		}
	}
	switch method {
	case "get":
		s.engine.GetWithStructParams(path, desc, dataReceiver, newHandler)
	case "post":
		s.engine.PostWithStructParams(path, desc, dataReceiver, newHandler)
	}
}
