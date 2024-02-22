/**************************************************
*Copyright(C).2016-2023,瀚辰光翼⽣物科技有限公司
*文件名：parse.go
*内容简述：*
*文件历史：
author 李承益 创建 2023/10/18
**************************************************/

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ReadAndParseConfig
// Description: 读取并解析配置
// Author: 李承益
// Date: 2023-09-05 10:10:43
// Param path string :
// Param ext interface{} :
// Return error :
func ReadAndParseConfig(path string, ext interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, ext)
}

// SaveToPath
// Description: 保存到指定路径
// Author: 李承益
// Date: 2023-10-18 11:26:16
// Param path string :
// Param ext interface{} :
// Return error :
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

// CreateFile
// Description: 创建文件
// Author: 李承益
// Date: 2023-08-16 16:00:34
// Param name string :
// Return f *os.File :
// Return err error :
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
