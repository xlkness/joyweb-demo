package controllers

import (
	"fmt"
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
)

func (ctl *Controllers) AddEnv(ctx *model.MyContext, params *request.EnvData) {
	storeData, err := ctl.Db.AddEnv(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(storeData)
}

func (ctl *Controllers) EditEnv(ctx *model.MyContext, params *request.EnvData) {
	fmt.Printf("receive edit env, data:%+v\n", params)

	storeData, err := ctl.Db.UpdateEnv(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(storeData)
}

func (ctl *Controllers) DeleteEnv(ctx *model.MyContext, params *request.DeleteEnv) {
	oldData, err := ctl.Db.DeleteEnv(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(oldData)
}
