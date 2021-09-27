package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string // absPath字段存放ZIP或JAR文件的绝对路径
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开zip文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	// 遍历文件 如果找到class则读取
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
