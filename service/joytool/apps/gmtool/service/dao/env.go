package dao

import (
	"fmt"
	"joytool/apps/gmtool/model/do"
	"joytool/apps/gmtool/model/request"
	"time"
)

func (d *Dao) QueryEnvList(project string) ([]*do.Env, error) {
	respEnvList := d.envList.GetList()
	var respEnvList1 []*do.Env
	for _, v := range respEnvList {
		if v.Data.Project == project {
			respEnvList1 = append(respEnvList1, v.Data)
		}
	}
	return respEnvList1, nil
}

func (d *Dao) AddEnv(env *request.EnvData) (*do.Env, error) {
	do := &do.Env{
		Project:   env.Project,
		Name:      env.Name,
		Desc:      env.Desc,
		Index:     env.Index,
		CreatedAt: time.Now().Format(time.DateTime),
	}
	_, find := d.envList.StoreUnique(do, do.Project, do.Name)
	if find {
		return nil, fmt.Errorf("存在同名:%v,%v", do.Project, do.Name)
	}
	return do, nil
}

func (d *Dao) UpdateEnv(env *request.EnvData) (*do.Env, error) {
	oldData, find := d.envList.Get(env.Project, env.Name)
	if !find {
		return nil, fmt.Errorf("数据库找不到数据:%v,%v", env.Project, env.Name)
	}
	oldData.Data.Desc = env.Desc
	oldData.Data.Index = env.Index
	d.envList.StoreOrReplace(oldData.Data, env.Project, env.Name)
	return oldData.Data, nil
}

func (d *Dao) DeleteEnv(env *request.DeleteEnv) (*do.Env, error) {
	data, find := d.envList.Delete(env.Project, env.Name)
	if !find {
		return nil, fmt.Errorf("数据库找不到数据:%v,%v", env.Project, env.Name)
	}

	return data.Data, nil
}
