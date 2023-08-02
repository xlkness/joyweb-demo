package dao

import (
	"fmt"
	"joytool/apps/gmtool/model/do"
	"joytool/apps/gmtool/model/request"
	"time"
)

func (d *Dao) PermissionGroupList(project string) []*do.PermissionGroupData {
	list := d.permissionGroupList.GetList()
	list1 := make([]*do.PermissionGroupData, 0, len(list))
	for _, v := range list {
		if v.Data.Project == project {
			list1 = append(list1, v.Data)
		}
	}
	return list1
}

func (d *Dao) GetPermissionGroup(project, group string) (*do.PermissionGroupData, bool) {
	if group == "" {
		return &do.PermissionGroupData{}, false
	}
	data, find := d.permissionGroupList.Get(project, group)
	if !find {
		return &do.PermissionGroupData{}, false
	}

	return data.Data, true
}

func (d *Dao) CreatePermissionGroup(params *request.PermissionGroupData) (*do.PermissionGroupData, error) {
	doData := &do.PermissionGroupData{
		Project:     params.Project,
		Name:        params.Name,
		Permissions: params.Permissions,
		CreatedAt:   time.Now().Format(time.DateTime),
	}

	_, find := d.permissionGroupList.StoreUnique(doData, params.Project, params.Name)
	if find {
		return nil, fmt.Errorf("数据库存在同名")
	}
	return doData, nil
}

func (d *Dao) EditPermissionGroup(params *request.PermissionGroupData) (*do.PermissionGroupData, error) {
	oldData, find := d.permissionGroupList.Get(params.Project, params.Name)
	if !find {
		return nil, fmt.Errorf("数据库没找到数据")
	}

	oldData.Data.Desc = params.Desc
	oldData.Data.Permissions = params.Permissions
	d.permissionGroupList.Update(oldData.Data, params.Name)

	return oldData.Data, nil
}

func (d *Dao) DeletePermissionGroup(params *request.DeletePermissionGroup) (*do.PermissionGroupData, error) {
	oldData, find := d.permissionGroupList.Delete(params.Project, params.Name)
	if !find {
		return nil, fmt.Errorf("数据库没找到数据")
	}
	return oldData.Data, nil
}
