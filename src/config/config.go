/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：config.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package config

import (
	"log"

	"tool/guitar/src/win"
)

type BaseConfigI interface {
	ReadConfig()
	SaveConfig()
}

var (
	configs []BaseConfigI
)

func AddConfig(config BaseConfigI) {
	config.ReadConfig()
	configs = append(configs, config)
}

func init() {
	win.AddExitFun(func() {
		log.Println("保存配置")
		for i := range configs {
			configs[i].SaveConfig()
		}
	})
}
