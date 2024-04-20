package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ReadAndParseConfig(path string, ext interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, ext)
}

func SaveToPath(path string, ext interface{}) error {
	marshal, err := json.Marshal(ext)
	if err != nil {
		return err
	}
	f, err := CreateFile(path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Truncate(0)
	_, err = f.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(name string) (f *os.File, err error) {
	path := filepath.Dir(name)
	if path != "." {
		// 检查文件夹是否创建
		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				// 文件夹不存在，创建文件夹
				err = os.MkdirAll(path, os.ModePerm)
				fmt.Println(err)
				if err != nil {
					return
				}
			} else {
				return
			}
		}
	}
	return os.Create(name)
}
