/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：exit.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package win

import (
	"log"
	"os"
	"os/signal"
)

func AddExitFun(f func()) {
	exitFun = append(exitFun, f)
}

var exitFun []func()

func init() {
	go func() {
		exitChan := make(chan os.Signal)
		signal.Notify(exitChan, os.Kill, os.Interrupt)
		select {
		case <-exitChan:
			log.Println("中断退出")
			exit()
		}

	}()
}

func exit() {
	for i := range exitFun {
		exitFun[i]()
	}
}
