package server

import (
	"joytool/apps/gmtool/model/request"
	"joytool/apps/gmtool/server/controllers"
)

func (s *Server) initRouter() {
	ctl := controllers.NewControllers(s.svc.Db)

	s.engine.GetWithStructParams("/gmtool", "gmtool", request.HomeView{}, ctl.HomeView)

	s.engine.PostWithStructParams("/addcommandserver", "添加指令服务器", request.CommandServerData{}, ctl.AddCommandServer)
	s.engine.PostWithStructParams("/editcommandserver", "编辑指令服务器", request.CommandServerData{}, ctl.EditCommandServer)
	s.engine.PostWithStructParams("/deletecommandserver", "编辑指令服务器", request.DeleteCommandServer{}, ctl.DeleteCommandServer)
	s.engine.PostWithStructParams("/deleteexechistory", "删除执行历史", request.DeleteExecHistory{}, ctl.DeleteExecHistory)

	s.engine.PostWithStructParams("/addenv", "添加环境", request.EnvData{}, ctl.AddEnv)
	s.engine.PostWithStructParams("/editenv", "编辑环境", request.EnvData{}, ctl.EditEnv)
	s.engine.PostWithStructParams("/deleteenv", "编辑环境", request.DeleteEnv{}, ctl.DeleteEnv)

	s.engine.GetWithStructParams("/commandlist", "get command list", request.CommandList{}, ctl.CommandList)
	s.engine.PostWithStructParams("/execcommand", "exec command list", request.ExecCommand{}, ctl.ExecCommand)

	s.engine.PostWithStructParams("/likecommand", "exec command list", request.ExecCommand{}, ctl.LikeCommand)
	s.engine.PostWithStructParams("/dislikecommand", "exec command list", request.DeleteLikeCommand{}, ctl.DislikeCommand)

	s.engine.GetWithStructParams("/listpermissiongroup", "exec command list", request.PermissionGroupList{}, ctl.PermissionGroupList)
	s.engine.PostWithStructParams("/createpermissiongroup", "exec command list", request.PermissionGroupData{}, ctl.CreatePermissionGroup)
	s.engine.PostWithStructParams("/editpermissiongroup", "exec command list", request.PermissionGroupData{}, ctl.EditPermissionGroup)
	s.engine.PostWithStructParams("/deletepermissiongroup", "exec command list", request.DeletePermissionGroup{}, ctl.DeletePermissionGroup)
}
