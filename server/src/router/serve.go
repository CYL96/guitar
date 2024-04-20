/**************************************************
*文件名：serve.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/27
**************************************************/

package hd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "server/docs"
	"server/src/config"
	"server/src/runCtx"
)

func Serve(ctx *runCtx.RunCtx) (err error) {
	host := fmt.Sprintf("%s:%d", config.SystemConfig.RunIp, config.SystemConfig.RunPort)
	r := gin.Default()
	r.Use(cors())

	GinRouter(r)
	// swag 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ctx.Info("swagger:", fmt.Sprintf("http://%s/swagger/index.html", host))

	r.Static("/Guitar", "Guitar")
	r.Static("/assets", "dist/assets")
	r.Static("/icon", "icon")
	r.StaticFile("/touch_001.mp3", "dist/touch_001.mp3")
	r.GET("/", func(c *gin.Context) {
		c.File("dist/index.html")
	})
	return r.Run(host)
}

// cors
// Description: 跨域配置
// Author: CYL96
// Date: 2023-08-28 14:28:56
// Return gin.HandlerFunc :
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		// c.Header("Access-Control-Allow-Headers", "token,Content-Type")
		// c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		// c.Set("content-type", "application/json")
		c.Next()
	}
}
