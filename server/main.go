package main

import (
	"os"
	"time"

	"server/src/common"
	"server/src/config"
	hd "server/src/router"
	"server/src/runCtx"
)

func main() {
	ctx := runCtx.DefaultContext()
	ctx.Info("程序开始启动...")

	ctx.Info("读取配置...")
	err := config.InitConfig(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		common.PrinlnIpRequest(ctx)
	}()

	if err != nil {
		return
	}
	ctx.Info("启动服务器配置...")

	err = hd.Serve(ctx)
	if err != nil {
		ctx.Error("启动服务器配置 err:", err)
		os.Exit(-1)
		return
	}

}
