package cwguitar

import (
	"log"

	"server/src/config"
)

const (
	configPath = "config/cwguitar.json"
)

var CWGuitarConfig = &CWGuitarConfigExt{}

type (
	CWGuitarConfigExt struct {
		Url   string `json:"url"`   // 访问链接
		Token string `json:"token"` // token
	}
)

func (C *CWGuitarConfigExt) ReadConfig() {
	log.Println("读取配置：", configPath)
	err := config.ReadAndParseConfig(configPath, C)
	if err != nil {
		log.Println(err)
	}
}

func (C *CWGuitarConfigExt) SaveConfig() {
	log.Println("保存配置：", configPath)
	err := config.SaveToPath(configPath, C)
	if err != nil {
		log.Println(err)
	}
}
