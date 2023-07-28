package controllers

import (
	"encoding/json"
	"fmt"
	"joytool/apps/gmtool/model"
	"joytool/apps/gmtool/model/do"
	"joytool/apps/gmtool/model/request"
	"time"
)

func (ctl *Controllers) CommandList(ctx *model.MyContext, req *request.CommandList) {
	fmt.Printf("receive command list\n")
	cmdServer, find, err := ctl.Db.QueryCmdServerByName(req.Project, req.Name)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	if !find {
		ctx.RespFailMessage(300, fmt.Sprintf("not found cmd info by id %v", req.Name))
		return
	}

	resp, err := queryCommandServerCommandList(cmdServer.Addr)
	if err != nil {
		fmt.Printf("queryCommandServerCommandList addr %v error:%v\n", cmdServer.Addr, err)
		ctx.RespFailMessage(300, err.Error())
		return
	}
	ctx.RespSuccessJson(resp)
}

func (ctl *Controllers) ExecCommand(ctx *model.MyContext, params *request.ExecCommand) {
	fmt.Printf("receive exec command:%+v\n", params)
	cmdServer, find, err := ctl.Db.QueryCmdServerByName(params.Project, params.CommandServerName)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	if !find {
		ctx.RespFailMessage(300, fmt.Sprintf("not found cmd info by id %v", params.CommandServerName))
		return
	}

	// 校验是否有这个指令
	rejectMessage := ""
	resp, err := queryCommandServerCommandList(cmdServer.Addr)
	for _, queryCmd := range resp {
		if queryCmd.Name == params.ExecBaseInfo.Name {
			// 找到gm指令，准备执行
			for _, execField := range params.ExecBaseInfo.Fields {
				var find = false
				for _, queryField := range queryCmd.Fields {
					if execField.Name == queryField.Name {
						// 校验参数是否存在，否则认为非法
						find = true
						break
					}
				}
				if !find {
					rejectMessage = fmt.Sprintf("指令[%v]没有在远程服务器当前描述信息中找到字段[%v]",
						params.ExecBaseInfo.Name, execField.Name)
					break
				}
			}

			if ok, err := ctl.execPermission(ctx, params.Project, params.Env, params.CommandServerName, queryCmd.Action); err != nil {
				ctx.RespFailMessage(300, err.Error())
				return
			} else if !ok {
				ctx.RespFailMessage(300, fmt.Sprintf("您无权执行这个命令！"))
				return
			}
		}
	}

	if rejectMessage != "" {
		ctx.RespFailMessage(300, rejectMessage)
		return
	}

	urlParams := ""
	for _, v := range params.ExecBaseInfo.Fields {
		urlParams += fmt.Sprintf("%v=%v&", v.Name, v.Value)
	}
	if len(params.ExecBaseInfo.Fields) > 0 {
		urlParams = urlParams[:len(urlParams)-1]
	}

	// 校验通过，准备执行指令
	res, code, err := execCommandServerCommand(cmdServer.Addr, params.ExecBaseInfo.Name, urlParams)

	// 记录执行历史
	resBin, _ := json.Marshal(res)
	record := &do.PostExecCommand{
		RequestInfo: &do.PreExecCommand{
			CommandBaseInfo:   params.ExecBaseInfo,
			Project:           params.Project,
			CommandServerName: params.CommandServerName,
			Env:               params.Env,
			User:              "test",
			Date:              time.Now().Format(time.DateTime),
		},
		ExecResCode: code,
		ExecRes:     string(resBin),
	}
	cmdServer.ExecHistoryList = append(cmdServer.ExecHistoryList, record)
	ctl.Db.SaveCommandServerList()

	// 响应前端
	if err != nil {
		//binData, _ := json.Marshal(res)
		ctx.RespFailMessage(code, string(err.Error()))
		return
	}
	if code != 200 {
		binData, _ := json.Marshal(res)
		ctx.RespFailMessage(code, string(binData))
		return
	}
	ctx.RespSuccessJson(res)
}
