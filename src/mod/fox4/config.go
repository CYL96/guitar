/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：config.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package fox4

import (
	"log"

	"tool/guitar/src/config"
)

const (
	configPath = "config/fox4.json"
)

var Fox4Config = &Fox4ConfigExt{}

type (
	Fox4ConfigExt struct {
		Url   string `json:"url"`   // 访问链接
		Token string `json:"token"` // token
	}
)

func (C *Fox4ConfigExt) ReadConfig() {
	log.Println("读取配置：", configPath)
	err := config.ReadAndParseConfig(configPath, C)
	if err != nil {
		log.Println(err)
	}
}

func (C *Fox4ConfigExt) SaveConfig() {
	log.Println("保存配置：", configPath)
	err := config.SaveToPath(configPath, C)
	if err != nil {
		log.Println(err)
	}
}
