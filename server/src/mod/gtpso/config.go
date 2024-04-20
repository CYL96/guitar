package gtpso

import (
	"log"

	"server/src/config"
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
