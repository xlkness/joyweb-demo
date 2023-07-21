package filedb

import (
	"fmt"
	"os"
)

var (
	FileDBDataPath = "db/"
)

type FileDB struct {
	SubPath  string
	FileName string
}

func NewFileDB(subPath, dbName string) (*FileDB, error) {
	fileDB := &FileDB{
		SubPath:  subPath,
		FileName: dbName,
	}

	fd, err := fileDB.getFd(false)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	return fileDB, nil
}

func (fdb *FileDB) GetCurrentObject() ([]byte, error) {
	fd, err := fdb.getFd(false)
	content, err := readFileAll(fd)
	if err != nil {
		return nil, fmt.Errorf("read file(%v) content error:%v", fdb.FileName, err)
	}
	defer fd.Close()
	return content, err
}

func (fdb *FileDB) WriteCurrentObject(data []byte) error {
	fd, err := fdb.getFd(true)
	if err != nil {
		return err
	}
	_, err = fd.Write(data)
	defer fd.Close()
	return err
}

func (fdb *FileDB) getFd(truncate bool) (*os.File, error) {
	os.MkdirAll(FileDBDataPath+fdb.SubPath, 0777)
	file := FileDBDataPath + fdb.SubPath + "/" + fdb.FileName + ".db"
	flags := os.O_CREATE | os.O_RDWR
	if truncate {
		flags |= os.O_TRUNC
	}
	fd, err := os.OpenFile(file, flags, 0777)
	if err != nil {
		return nil, fmt.Errorf("open file(%v) error:%v", file, err)
	}

	return fd, nil
}
