package fox4

import (
	"log"

	"server/src/config"
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
