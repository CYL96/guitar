/**************************************************
*文件名：system.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"server/src/runCtx"
)

type (
	SystemExt struct {
		RunIp   string `json:"run_ip" default:"" example:""`           // 运行ip地址
		RunPort int    `json:"run_port" default:"8080" example:"8080"` // 运行端口
	}
)

var SystemConfig SystemExt
var systemLk sync.RWMutex

func readSystemConfig(ctx *runCtx.RunCtx) (err error) {
	file, err := os.ReadFile("./config/system.json")
	if err != nil {
		ctx.Warn(err)
		CreateNormalSystemConfig(ctx)
		err = nil
		return
	}
	err = json.Unmarshal(file, &SystemConfig)
	if err != nil {
		return
	}

	return nil
}

func saveSystemConfig(ctx *runCtx.RunCtx) (err error) {
	create, err := os.Create("./config/system.json")
	if err != nil {
		ctx.Warn(err)
		return
	}
	defer create.Close()
	marshal, err := json.Marshal(SystemConfig)
	if err != nil {
		ctx.Warn(err)
		return
	}
	_, err = create.Write(marshal)
	if err != nil {
		ctx.Warn(err)
		return
	}
	return nil
}

func WriteConfigToWev(ctx *runCtx.RunCtx) (err error) {
	var (
		create *os.File
	)

	create, err = os.Create("./dist/config.json")
	if err != nil {
		ctx.Error(err)
		return
	}
	defer create.Close()
	config := make(map[string]string)
	config["server"] = fmt.Sprintf("http://%s:%d", SystemConfig.RunIp, SystemConfig.RunPort)
	marshal, err := json.Marshal(config)
	if err != nil {
		ctx.Error(err)
		return
	}
	_, err = create.Write(marshal)
	if err != nil {
		ctx.Error(err)
		return
	}
	return nil

}

func CreateNormalSystemConfig(ctx *runCtx.RunCtx) {
	create, err := os.Create("./config/system.json")
	if err != nil {
		ctx.Warn(err)
		return
	}
	defer create.Close()
	normal := SystemExt{
		RunIp:   "0.0.0.0",
		RunPort: 55001,
	}
	marshal, err := json.Marshal(normal)
	if err != nil {
		ctx.Warn(err)
		return
	}
	_, err = create.Write(marshal)
	if err != nil {
		ctx.Warn(err)
		return
	}
}

func GetSystemConfig(ctx *runCtx.RunCtx) GetSystemConfigResult {
	systemLk.RLock()
	defer systemLk.RUnlock()
	return GetSystemConfigResult{SystemConfig}
}

type (
	GetSystemConfigResult struct {
		SystemExt
	}
)

func UpdateSystemConfig(ctx *runCtx.RunCtx, para UpdateSystemConfigPara) (err error) {
	systemLk.RLock()
	defer systemLk.RUnlock()
	SystemConfig.RunPort = para.RunPort

	return saveSystemConfig(ctx)
}

type (
	UpdateSystemConfigPara struct {
		SystemExt
	}
)
