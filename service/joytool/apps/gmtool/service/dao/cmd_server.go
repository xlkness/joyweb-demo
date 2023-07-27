package dao

import (
	"fmt"
	"joytool/apps/gmtool/model/do"
	"joytool/apps/gmtool/model/request"
	"time"
)

func (d *Dao) QueryCmdServerList(project string) ([]*do.CommandServer, error) {
	respCmdServerList := d.cmdServerList.GetList()
	var respCmdServerList1 []*do.CommandServer
	for _, v := range respCmdServerList {
		if v.Data.Project == project {
			respCmdServerList1 = append(respCmdServerList1, v.Data)
		}
	}

	//commandServerList := []*model.CommandServer{
	//	{"test1", "渠道指令服务器", "192.168.1.104:8888", []string{"pro", "test"}, time.Now(), []string{}, nil},
	//	{"test2", "线上指令服务器", "192.168.1.104:8888", []string{"pro"}, time.Now(), []string{}, nil},
	//}
	//
	//envs := []*model.Env{
	//	{"pro", "线上环境", 1},
	//	{"test", "测试环境", 2},
	//}

	return respCmdServerList1, nil
}

func (d *Dao) QueryCmdServerByName(project, name string) (*do.CommandServer, bool, error) {
	cmdServer, find := d.cmdServerList.Get(project, name)
	if !find {
		return nil, false, nil
	}
	return cmdServer.Data, find, nil
	//cmd := &model.CommandServer{"test1", "渠道指令服务器", "192.168.1.104:8888", []string{"pro", "test"}, time.Now(), []string{}, nil}

	//return cmd, true, nil
}

func (d *Dao) AddCommandServer(cmdServer *request.CommandServerData) (*do.CommandServer, error) {
	do := &do.CommandServer{
		Project:         cmdServer.Project,
		Name:            cmdServer.Name,
		Desc:            cmdServer.Desc,
		Addr:            cmdServer.Addr,
		Envs:            cmdServer.Envs,
		CreatedAt:       time.Now().Format(time.DateTime),
		ExecHistoryList: nil,
	}
	_, find := d.cmdServerList.StoreUnique(do, cmdServer.Project, cmdServer.Name)
	if find {
		return nil, fmt.Errorf("数据库存在同名:%v", cmdServer.Name)
	}
	return do, nil
}

func (d *Dao) SaveCommandServerList() error {
	return d.cmdServerList.PersistentSave()
}

func (d *Dao) UpdateCommandServer(cmdServer *request.CommandServerData) (*do.CommandServer, error) {
	oldData, find := d.cmdServerList.Get(cmdServer.Project, cmdServer.Name)
	if !find {
		return nil, fmt.Errorf("数据库找不到数据:%v,%v", cmdServer.Project, cmdServer.Name)
	}
	if oldData.Data.Project != cmdServer.Project {
		return nil, fmt.Errorf("不能修改项目名！%v/%v", oldData.Data.Project, cmdServer.Project)
	}

	oldData.Data.Desc = cmdServer.Desc
	oldData.Data.Addr = cmdServer.Addr
	oldData.Data.Envs = cmdServer.Envs
	d.cmdServerList.Update(oldData.Data, cmdServer.Project, cmdServer.Name)
	return oldData.Data, nil
}

func (d *Dao) DeleteCommandServer(cmdServer *request.DeleteCommandServer) (*do.CommandServer, error) {
	data, find := d.cmdServerList.Delete(cmdServer.Project, cmdServer.Name)
	if !find {
		return nil, fmt.Errorf("数据库找不到数据:%v,%v", cmdServer.Project, cmdServer.Name)
	}

	return data.Data, nil
}

func (d *Dao) DeleteCommandServerExecHistory(params *request.DeleteExecHistory) error {
	data, find := d.cmdServerList.Get(params.Project, params.Name)
	if !find {
		return fmt.Errorf("数据库找不到数据:%v,%v", params.Project, params.Name)
	}

	if len(data.Data.ExecHistoryList) <= params.Index {
		return fmt.Errorf("删除的索引不合法:%v,%v,%v/%v", params.Project, params.Name, params.Index, len(data.Data.ExecHistoryList))
	}

	data.Data.ExecHistoryList = append(data.Data.ExecHistoryList[:params.Index],
		data.Data.ExecHistoryList[params.Index+1:]...)

	d.cmdServerList.PersistentSave()

	return nil
}
