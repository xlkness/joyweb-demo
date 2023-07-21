package server

func (s *Server) initRouter() {
	s.engine.PostWithStructParams("/authorize", "授权", nil, Authorize)
	s.engine.Get("/authorize", "授权", Authorize)
	s.engine.PostWithStructParams("/token", "token", nil, Token)
	s.engine.Get("/token", "token", Token)
	s.engine.PostWithStructParams("/login", "登陆", LoginData{}, Login)
}
