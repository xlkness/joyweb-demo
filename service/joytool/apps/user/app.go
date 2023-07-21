package user

import (
	lkit_go "github.com/xlkness/lkit-go"
	"joytool/apps/user/model"
	"joytool/apps/user/server"
	"joytool/apps/user/service"
)

var bootConf = &model.BootConfFlag{}

func NewApp() *lkit_go.ApplicationDescInfo {
	app := lkit_go.NewApplicationDescInfo("user", func(globalBootFlag *lkit_go.CommBootFlag, _ interface{}, app *lkit_go.Application) error {
		svc := service.NewService()
		httpEngine := lkit_go.NewEngine(":"+bootConf.ServerPort, func() lkit_go.WebEngineContext {
			return new(model.MyContext)
		})
		server.NewServer(httpEngine, svc)
		app.WithServer("gm管理系统服务器", httpEngine)
		return nil
	}, lkit_go.WithAppBootFlag(bootConf))
	return app
}
