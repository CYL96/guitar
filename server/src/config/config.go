package config

import "server/src/runCtx"

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

func InitConfig(ctx *runCtx.RunCtx) (err error) {

	ctx.Info("读取系统配置")
	err = readSystemConfig(ctx)
	if err != nil {
		ctx.Error("读取系统配置 err:", err)
		return
	}
	return

}
