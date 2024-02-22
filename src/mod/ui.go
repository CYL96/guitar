/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：ui.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package mod

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"tool/guitar/src/mod/cwguitar"
	"tool/guitar/src/mod/fox4"
	"tool/guitar/src/mod/gtpso"
)

type (
	ListViewExt struct {
		ShowName string
		ViewUi   fyne.CanvasObject
	}
)

var (
	listViewList = make([]ListViewExt, 0)
)

func AddView(name string, view fyne.CanvasObject) {
	listViewList = append(listViewList, ListViewExt{
		ShowName: name,
		ViewUi:   view,
	})
}

func addView() {
	AddView(cwguitar.CwguitarName(), cwguitar.CwguitarView())
	AddView(gtpso.GtpsoName(), gtpso.GtpsoView())
	AddView(fox4.Fox4Name(), fox4.Fox4View())
	AddView("test1", nil)

}

func InitMainView() fyne.CanvasObject {
	sync.OnceFunc(addView)()
	editView := new(fyne.Container)
	list := widget.NewList(
		func() int {
			return len(listViewList)
		},
		func() fyne.CanvasObject {
			ctr := container.NewHBox(widget.NewIcon(theme.FileApplicationIcon()), widget.NewLabel("Template Object"))
			return ctr
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(listViewList[id].ShowName)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		editView.RemoveAll()
		editView.Add(listViewList[id].ViewUi)
		editView.Refresh()
	}
	if len(listViewList) > 0 {
		list.Select(0)
	}
	// hSplit := container.NewHSplit(list, editView)
	// hSplit.Offset = -100
	// return container.NewBorder(nil, nil, list, nil, editView)
	return container.NewHBox(list, editView)
}
