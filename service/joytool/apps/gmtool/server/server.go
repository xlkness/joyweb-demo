package server

import (
	lkit_go "github.com/xlkness/lkit-go"
	"joytool/apps/gmtool/server/controllers"
	"joytool/apps/gmtool/service"
)

type Server struct {
	engine *lkit_go.WebEngine
	ctl    *controllers.Controllers
	svc    *service.Service
}

func NewServer(engine *lkit_go.WebEngine, svc *service.Service) *Server {
	s := new(Server)
	s.engine = engine
	s.svc = svc

	s.engine.Use(jwtMiddleWare())
	s.engine.GetGinEngine().Use(filtterWare()).Use(Cors())
	s.initRouter()

	return s
}
