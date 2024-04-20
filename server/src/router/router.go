/**************************************************
*文件名：router.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package hd

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GinRouter(e *gin.Engine) {
	api := e.Group("/api")
	api.POST("/Exit", func(c *gin.Context) {
		os.Exit(0)
	})
	api.POST("/GetGuitarModList", GetGuitarModList)
	api.POST("/GetGuitarModInfo", GetGuitarModInfo)
	api.POST("/DownloadGuitarModInfo", DownloadGuitarModInfo)
	api.POST("/DownloadGuitar", DownloadGuitar)
	api.POST("/GetGuitarPicList", GetGuitarPicList)
	api.POST("/DeleteGuitarClass", DeleteGuitarClass)
	api.POST("/RenameGuitarClass", RenameGuitarClass)
}
