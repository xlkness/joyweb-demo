package dao

import (
	"fmt"
	doUser "joytool/apps/user/model/do"
	"joytool/apps/user/model/request"
	"time"
)

func (d *Dao) GetUserInfo(userName, system string) (*doUser.User, bool) {
	dbData, find := d.userList.Get(userName)
	if !find {
		return nil, find
	}

	curSystemData := &doUser.SystemUserGroup{
		Name: system,
	}
	for _, v := range dbData.Data.Systems {
		if v.Name == system {
			curSystemData.IsAdmin = v.IsAdmin
			curSystemData.Group = v.Group
			break
		}
	}

	if userName == "admin" {
		curSystemData.IsAdmin = true
	}

	newUserData := dbData.Data.Copy()
	newUserData.Systems = []*doUser.SystemUserGroup{curSystemData}
	return newUserData, find
}

func (d *Dao) UserList() ([]*doUser.User, error) {
	list := d.userList.GetList()
	list1 := make([]*doUser.User, 0, len(list))
	for _, v := range list {
		if v.Data.UserName == "admin" {
			continue
		}
		list1 = append(list1, v.Data)
	}
	return list1, nil
}

func (d *Dao) CreateUser(userData *request.CreateUser) (*doUser.User, error) {
	do := &doUser.User{
		Systems:   userData.GroupList,
		UserName:  userData.BaseInfo.UserName,
		Password:  userData.BaseInfo.PassWord,
		Salt:      "sfdsd",
		CreatedAt: time.Now().Format(time.DateTime),
	}

	_, find := d.userList.StoreUnique(do, userData.BaseInfo.UserName)
	if find {
		return nil, fmt.Errorf("存在同名用户！")
	}

	return do, nil
}

func (d *Dao) EditUser(newData *request.CreateUser) (*doUser.User, error) {
	oldData, find := d.userList.Get(newData.BaseInfo.UserName)
	if !find {
		return nil, fmt.Errorf("数据库没有找到数据")
	}

	if len(newData.GroupList) > 0 {
		for _, v := range oldData.Data.Systems {
			if v.Name == newData.GroupList[0].Name {
				v.IsAdmin = newData.GroupList[0].IsAdmin
				v.Group = newData.GroupList[0].Group
				goto Next
			}
		}

		oldData.Data.Systems = append(oldData.Data.Systems, newData.GroupList[0])
	}

Next:
	oldData, _ = d.userList.Update(oldData.Data, oldData.Data.UserName)

	return oldData.Data, nil
}

func (d *Dao) DeleteUser(userData *request.DeleteUser) (*doUser.User, error) {
	dbData, find := d.userList.Delete(userData.UserName)
	if !find {
		return nil, fmt.Errorf("数据库没有找到数据")
	}

	return dbData.Data, nil
}

func (d *Dao) LoginVerify(userData *request.LoginData) (*doUser.User, bool, error) {
	do, find := d.userList.Get(userData.BaseInfo.UserName)
	if !find {
		return nil, false, fmt.Errorf("数据库没有找到用户！")
	}

	if do.Data.Password != userData.BaseInfo.PassWord {
		return nil, false, nil
	}

	return do.Data, true, nil
}
