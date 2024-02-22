/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：config.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package gtpso

import (
	"log"

	"tool/guitar/src/config"
)

const (
	configPath = "config/gtpso.json"
)

var GtpsoConfig = &GtpsoConfigExt{}

type (
	GtpsoConfigExt struct {
		Url   string `json:"url"`   // 访问链接
		Token string `json:"token"` // token
	}
)

func (C *GtpsoConfigExt) ReadConfig() {
	log.Println("读取配置：", configPath)
	err := config.ReadAndParseConfig(configPath, C)
	if err != nil {
		log.Println(err)
	}
}

func (C *GtpsoConfigExt) SaveConfig() {
	log.Println("保存配置：", configPath)
	err := config.SaveToPath(configPath, C)
	if err != nil {
		log.Println(err)
	}
}
