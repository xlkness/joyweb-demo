package dao

import (
	"encoding/json"
	"fmt"
	"joytool/apps/gmtool/model/do"
	"joytool/lib/filedb"
	"sort"
	"strconv"
	"sync"
)

type DbRowData[Elem do.Env | do.CommandServer | do.LikeExecCommand] struct {
	ID   int
	Data *Elem
}
type fileDb[Elem do.Env | do.CommandServer | do.LikeExecCommand] struct {
	db      *filedb.FileDB
	allData map[string]*DbRowData[Elem]
	maxID   int
	locker  *sync.Mutex
}

func newFileDb[Elem do.Env | do.CommandServer | do.LikeExecCommand](dbName string) (*fileDb[Elem], error) {
	db, err := filedb.NewFileDB("gmtool", dbName)
	if err != nil {
		return nil, err
	}

	allJsonData, err := db.GetCurrentObject()
	if err != nil {
		return nil, err
	}

	allData := make(map[string]*DbRowData[Elem])

	if len(allJsonData) != 0 {
		err = json.Unmarshal(allJsonData, &allData)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal error:%v", err)
		}
	}

	maxID := 0
	for _, v := range allData {
		if maxID < v.ID {
			maxID = v.ID
		}
	}

	fd := &fileDb[Elem]{
		db:      db,
		allData: allData,
		maxID:   maxID,
		locker:  new(sync.Mutex),
	}
	return fd, nil
}

func (fd *fileDb[Elem]) GetList() []*DbRowData[Elem] {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	list := make([]*DbRowData[Elem], 0, len(fd.allData))
	for _, v := range fd.allData {
		list = append(list, v)
	}

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	return list
}

func (fd *fileDb[Elem]) Get(unionName ...any) (*DbRowData[Elem], bool) {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	name := fd.joinName(unionName...)
	data, find := fd.allData[name]
	return data, find
}

// StoreUnique 不能重名
func (fd *fileDb[Elem]) StoreUnique(data *Elem, unionName ...any) (*DbRowData[Elem], bool) {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	name := fd.joinName(unionName...)
	_, find := fd.allData[name]
	if find {
		return nil, true
	}
	return fd.storeOrReplace(data, unionName...)
}

// StoreOrReplace 重名就替换
func (fd *fileDb[Elem]) StoreOrReplace(data *Elem, unionName ...any) (*DbRowData[Elem], bool) {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	return fd.storeOrReplace(data, unionName...)
}
func (fd *fileDb[Elem]) storeOrReplace(data *Elem, unionName ...any) (*DbRowData[Elem], bool) {
	name := fd.joinName(unionName...)
	oldData, find := fd.allData[name]
	if find {
		oldData.Data = data
	} else {
		oldData = &DbRowData[Elem]{
			ID:   fd.maxID + 1,
			Data: data,
		}
		fd.maxID += 1
	}

	fd.allData[name] = oldData
	fd.persistentSave()
	return oldData, find
}

// StorePush 直接追加数据，key为id
func (fd *fileDb[Elem]) StorePush(data *Elem) *DbRowData[Elem] {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	name := strconv.Itoa(fd.maxID + 1)
	storeData := &DbRowData[Elem]{
		ID:   fd.maxID + 1,
		Data: data,
	}
	fd.allData[name] = storeData
	fd.maxID += 1
	fd.persistentSave()
	return storeData
}

func (fd *fileDb[Elem]) Update(data *Elem, unionName ...any) (*DbRowData[Elem], bool) {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	name := fd.joinName(unionName...)
	storeData, find := fd.allData[name]
	if !find {
		return nil, false
	}
	storeData.Data = data
	fd.persistentSave()
	return storeData, true
}

func (fd *fileDb[Elem]) Delete(unionName ...any) (*DbRowData[Elem], bool) {
	fd.locker.Lock()
	defer fd.locker.Unlock()

	name := fd.joinName(unionName...)
	storeData, find := fd.allData[name]
	if !find {
		return nil, false
	}
	delete(fd.allData, name)
	fd.persistentSave()
	return storeData, true
}

func (fd *fileDb[Elem]) persistentSave() error {
	binData, err := json.Marshal(fd.allData)
	if err != nil {
		fmt.Printf("persistentSave error:%v, %v\n", fd.db.FileName, fd.allData)
		return fmt.Errorf("json marshal error:%v", err)
	}

	return fd.db.WriteCurrentObject(binData)
}

func (fd *fileDb[Elem]) joinName(unionName ...any) string {
	name := ""
	for _, v := range unionName {
		name += fmt.Sprintf("%v-", v)
	}
	name = name[:len(name)-1]
	return name
}
