/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：font.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package win

import (
	"log"
	"os"
)

func InitEnv() (err error) {
	var fyneFontKey = "FYNE_FONT"
	// 将环境变量 "FYNE_FONT" 设置为字体文件的路径
	err = os.Setenv(fyneFontKey, "./config/wryh.ttf")
	// 如果出现错误，打印错误并从函数返回
	if err != nil {
		return
	}
	AddExitFun(func() {
		log.Println("清除 Env:", fyneFontKey)
		os.Unsetenv(fyneFontKey)
	})
	return err
}
