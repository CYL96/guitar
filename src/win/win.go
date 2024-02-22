/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：win.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package win

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

var (
	App    fyne.App
	Window fyne.Window
)

func InitWin(title string) (err error) {
	// 创建一个新的应用
	App = app.New()
	// 设置应用的图标为设置图标
	App.SetIcon(theme.SettingsIcon())
	// 创建一个新的窗口，标题为 "配置"
	Window = App.NewWindow(title)
	Window.SetOnClosed(exit)
	return err
}
