package dao

import (
	"fmt"
	do "joytool/apps/gmtool/model/do"
	"joytool/apps/gmtool/model/request"
	"time"
)

func (d *Dao) LikeCommand(params *request.ExecCommand) (*do.LikeExecCommand, int, error) {
	do := &do.LikeExecCommand{
		Project: params.Project,
		User:    "test",
		Command: &do.PreExecCommand{
			CommandBaseInfo:   params.ExecBaseInfo,
			Project:           params.Project,
			CommandServerName: params.CommandServerName,
			Env:               params.Env,
			User:              "test",
			Date:              time.Now().Format(time.DateTime),
		},
	}
	fmt.Printf("like command do :%+v|%+v\n", do, do.Command)
	storeData := d.likeList.StorePush(do)
	do.ID = storeData.ID
	return storeData.Data, storeData.ID, nil
}

func (d *Dao) DislikeCommand(likeCommand *request.DeleteLikeCommand) (*do.LikeExecCommand, bool) {
	data, find := d.likeList.Delete(likeCommand.ID)
	if find {
		return data.Data, true
	}
	return nil, false
}

func (d *Dao) LikeCommandList(project, user string) ([]*do.LikeExecCommand, error) {
	list := d.likeList.GetList()
	var list1 []*do.LikeExecCommand
	for _, v := range list {
		v.Data.ID = v.ID
		if v.Data.Project == project && v.Data.User == user {
			list1 = append(list1, v.Data)
		}
	}
	return list1, nil
}
