/**************************************************
*文件名：log.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package runCtx

import (
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func getCaller(skip int) (call string) {
	skip++
	_, path, line, _ := runtime.Caller(skip)
	path = filepath.Base(path)
	return path + ":" + strconv.Itoa(line)
}

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
