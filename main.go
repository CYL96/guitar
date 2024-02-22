/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：main.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package main

import (
	"log"

	"fyne.io/fyne/v2"
	"tool/guitar/src/mod"
	"tool/guitar/src/win"
)

func main() {
	err := win.InitEnv()
	if err != nil {
		log.Println(err)
	}
	win.InitWin("test")

	win.Window.SetContent(mod.InitMainView())
	win.Window.Resize(fyne.NewSize(800, 600))

	win.Window.SetMaster()

	win.Window.ShowAndRun()
}
