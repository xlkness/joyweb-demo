package server

import "joytool/apps/user/model/request"

func (s *Server) initRouter() {
	s.engine.PostWithStructParams("/createuser", "登陆", request.CreateUser{}, s.CreateUser)
	s.engine.PostWithStructParams("/edituser", "登陆", request.CreateUser{}, s.EditUser)
	s.engine.PostWithStructParams("/deleteuser", "登陆", request.DeleteUser{}, s.DeleteUser)
	s.engine.PostWithStructParams("/login", "登陆", request.LoginData{}, s.Login)
	s.engine.Get("/logout", "登陆", s.Logout)
	s.engine.Get("/listusers", "登陆", s.ListUsers)

}
