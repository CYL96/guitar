/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：alter.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package win

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func Alter(title, content string) {
	log.Println("dialog")
	infoDialog := dialog.NewInformation(title, content, Window)
	infoDialog.Resize(fyne.NewSize(400, 300))
	infoDialog.Show()
}
