package common

import (
	"fmt"
	"net"

	"server/src/config"
	"server/src/runCtx"
)

func GetMyIpAddr(ctx *runCtx.RunCtx) (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		ctx.Error(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return
}

func PrinlnIpRequest(ctx *runCtx.RunCtx) {

	if config.SystemConfig.RunIp == "" ||
		config.SystemConfig.RunIp == "0.0.0.0" {
		ips := GetMyIpAddr(ctx)
		for _, ip := range ips {
			web := fmt.Sprintf("http://%s:%d/", ip, config.SystemConfig.RunPort)
			ctx.Info("访问链接:", web)
		}
	} else {
		web := fmt.Sprintf("http://%s:%d", config.SystemConfig.RunIp, config.SystemConfig.RunPort)
		ctx.Info("访问链接:", web)
	}

}
