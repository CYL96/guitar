package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	PATH = "Guitar"
)

func OpenRelativeExplorer(name string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	absDir := dir + "\\" + PATH + "\\" + name
	log.Println(absDir)
	exec.Command(`cmd`, `/c`, `explorer`, absDir).CombinedOutput()
}

func DownloadPic(name string, pics []string) (err error) {
	var (
		res *http.Response
	)
	strings.Replace(name, "/", "-", -1)
	strings.Replace(name, "\\", "-", -1)

	dir := PATH + "/" + name
	for i, v := range pics {
		filesuffix := path.Ext(v)
		res, err = http.Get(v)
		if err != nil {
			return
		}
		defer res.Body.Close()

		if 0 != len(dir) {
			is_exist := false
			is_exist, err = PathExists(dir)
			if nil != err {
				return
			}
			if is_exist == false {
				// 不存在文件夹时 先创建文件夹再上传
				err = os.MkdirAll(dir, os.ModePerm) // 创建文件夹
				if err != nil {
					return
				}
			}
		}
		storeName := fmt.Sprintf("%d%s", i+1, filesuffix)
		f, err := os.OpenFile(dir+"/"+storeName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		// 写入文件数据;
		io.Copy(f, res.Body)
		f.Close()
	}
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
