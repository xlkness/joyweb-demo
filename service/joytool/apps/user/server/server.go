package server

import (
	lkit_go "github.com/xlkness/lkit-go"
	"joytool/apps/user/service"
)

type Server struct {
	engine *lkit_go.WebEngine
	svc    *service.Service
}

func NewServer(engine *lkit_go.WebEngine, svc *service.Service) *Server {
	s := new(Server)
	s.engine = engine
	s.svc = svc

	s.engine.GetGinEngine().Use(jwtMiddleWare()).Use(Cors())
	s.initRouter()

	return s
}
