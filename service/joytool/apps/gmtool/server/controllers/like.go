package controllers

import (
	"fmt"
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/request"
)

// LikeCommand 收藏某个指令
func (ctl *Controllers) LikeCommand(ctx *model.MyContext, params *request.ExecCommand) {
	fmt.Printf("receive like:%+v\n", params)
	do, _, err := ctl.Db.LikeCommand(params)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(do)
}

func (ctl *Controllers) DislikeCommand(ctx *model.MyContext, params *request.DeleteLikeCommand) {
	do, find := ctl.Db.DislikeCommand(params)
	if find {
		ctx.RespSuccessJson(do)
	} else {
		ctx.RespFailMessage(300, fmt.Sprintf("数据库没找到数据:%v,%v", params.Project, params.ID))
	}
}
