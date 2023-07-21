package controllers

import (
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
)

func (ctl *Controllers) AddCommandServer(ctx *model.MyContext, params *request.CommandServerData) {
	storeData, err := ctl.Db.AddCommandServer(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(storeData)
}

func (ctl *Controllers) EditCommandServer(ctx *model.MyContext, params *request.CommandServerData) {
	storeData, err := ctl.Db.UpdateCommandServer(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(storeData)
}

func (ctl *Controllers) DeleteCommandServer(ctx *model.MyContext, params *request.DeleteCommandServer) {
	oldData, err := ctl.Db.DeleteCommandServer(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(oldData)
}

func (ctl *Controllers) DeleteExecHistory(ctx *model.MyContext, params *request.DeleteExecHistory) {
	err := ctl.Db.DeleteCommandServerExecHistory(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessMessage("ok")
}
