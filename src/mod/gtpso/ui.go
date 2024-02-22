/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：ui.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package gtpso

import (
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	homePage = "https://www.gtpso.com"
)

func init() {
	// config.AddConfig(GtpsoConfig)
}

func GtpsoName() string {
	return "Gtpso"
}
func GtpsoView() fyne.CanvasObject {
	ctr := container.NewVBox()
	toUrl, err := url.Parse(homePage)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	btn := widget.NewButton("open", func() {
		fyne.CurrentApp().OpenURL(toUrl)
	})
	ctr.Add(btn)

	ctr.Add(widget.NewLabel("url:"))
	var downUrl string
	ctr.Add(widget.NewEntryWithData(binding.BindString(&downUrl)))
	ctr.Add(widget.NewLabel("FileName:(Auto fill When Empty)"))

	var downName string
	ctr.Add(widget.NewEntryWithData(binding.BindString(&downName)))

	var (
		btn1, btn2 *widget.Button
	)
	btn1 = widget.NewButton("Download", func() {
		btn1.Disable()
		downLoad(false, downName, downUrl)
		btn1.Enable()
	})
	ctr.Add(btn1)

	btn2 = widget.NewButton("Download & Open", func() {
		btn2.Disable()
		downLoad(true, downName, downUrl)
		btn2.Enable()
	})
	ctr.Add(btn2)

	ctr.Resize(fyne.NewSize(600, 600))
	return ctr
}
