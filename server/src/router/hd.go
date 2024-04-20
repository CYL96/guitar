/**************************************************
*文件名：hd.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package hd

import (
	"github.com/gin-gonic/gin"
	"server/src/common"
	"server/src/mod"
	"server/src/mod/intf"
	"server/src/pic"
	"server/src/runCtx"
)

// GetGuitarModList
// @Summary	获取模块列表
// @Accept        json
// @Produce       json
// @Description	获取模块列表
// @Tags			Control
// @Param body body GetGuitarModListReq true "请求"
// @success 200 {object} GinResponse{data=GetGuitarModListResp} "desc"
// @Router			/api/GetGuitarModList [post]
func GetGuitarModList(c *gin.Context) {
	var err error
	var result GetGuitarModListResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetGuitarModListReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.List = mod.GetModList(ctx, req.GetModListPara)
	return
}

type (
	GetGuitarModListReq struct {
		mod.GetModListPara
	}
	GetGuitarModListResp struct {
		List []intf.GuitarModT `json:"list"`
	}
)

// GetGuitarModInfo
// @Summary	解析链接
// @Accept        json
// @Produce       json
// @Description	解析链接
// @Tags			Control
// @Param body body GetGuitarModInfoReq true "请求"
// @success 200 {object} GinResponse{data=GetGuitarModInfoResp} "desc"
// @Router			/api/GetGuitarModInfo [post]
func GetGuitarModInfo(c *gin.Context) {
	var err error
	var result GetGuitarModInfoResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetGuitarModInfoReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.GuitarGetInfo, err = mod.GetModGuitarInfo(ctx, req.GetModGuitarInfoPara)
	return
}

type (
	GetGuitarModInfoReq struct {
		mod.GetModGuitarInfoPara
	}
	GetGuitarModInfoResp struct {
		intf.GuitarGetInfo
	}
)

// DownloadGuitarModInfo
// @Summary	下载图片
// @Accept        json
// @Produce       json
// @Description	下载图片
// @Tags			Control
// @Param body body DownloadGuitarModInfoReq true "请求"
// @success 200 {object} GinResponse{data=DownloadGuitarModInfoResp} "desc"
// @Router			/api/DownloadGuitarModInfo [post]
func DownloadGuitarModInfo(c *gin.Context) {
	var err error
	var result DownloadGuitarModInfoResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req DownloadGuitarModInfoReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.DownloadModGuitarInfo(ctx, req.DownloadModGuitarInfoPara)
	return
}

type (
	DownloadGuitarModInfoReq struct {
		mod.DownloadModGuitarInfoPara
	}
	DownloadGuitarModInfoResp struct {
	}
)

// DownloadGuitar
// @Summary	下载图片
// @Accept        json
// @Produce       json
// @Description	下载图片
// @Tags			Control
// @Param body body DownloadGuitarReq true "请求"
// @success 200 {object} GinResponse{data=DownloadGuitarResp} "desc"
// @Router			/api/DownloadGuitar [post]
func DownloadGuitar(c *gin.Context) {
	var err error
	var result DownloadGuitarResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req DownloadGuitarReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = common.DownloadPic(req.Name, req.PicList)
	return
}

type (
	DownloadGuitarReq struct {
		intf.GuitarGetInfo
	}
	DownloadGuitarResp struct {
	}
)

// GetGuitarPicList
// @Summary	获取图片列表
// @Accept        json
// @Produce       json
// @Description	获取图片列表
// @Tags			Control
// @Param body body GetGuitarPicListReq true "请求"
// @success 200 {object} GinResponse{data=GetGuitarPicListResp} "desc"
// @Router			/api/GetGuitarPicList [post]
func GetGuitarPicList(c *gin.Context) {
	var err error
	var result GetGuitarPicListResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetGuitarPicListReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.List, err = pic.GetGuitarClassList(ctx, req.GetGuitarClassListPara)
	return
}

type (
	GetGuitarPicListReq struct {
		pic.GetGuitarClassListPara
	}
	GetGuitarPicListResp struct {
		List []pic.GuitarClassT `json:"list"`
	}
)

// DeleteGuitarClass
// @Summary	删除
// @Accept        json
// @Produce       json
// @Description	删除
// @Tags			Control
// @Param body body DeleteGuitarClassReq true "请求"
// @success 200 {object} GinResponse{} "desc"
// @Router			/api/DeleteGuitarClass [post]
func DeleteGuitarClass(c *gin.Context) {
	var err error
	var result DeleteGuitarClassResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req DeleteGuitarClassReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = pic.DeleteGuitarClass(ctx, req.DeleteGuitarClassPara)
	return
}

type (
	DeleteGuitarClassReq struct {
		pic.DeleteGuitarClassPara
	}
	DeleteGuitarClassResp struct {
	}
)

// RenameGuitarClass
// @Summary	修改名称
// @Accept        json
// @Produce       json
// @Description	修改名称
// @Tags			Control
// @Param body body RenameGuitarClassReq true "请求"
// @success 200 {object} GinResponse{} "desc"
// @Router			/api/RenameGuitarClass [post]
func RenameGuitarClass(c *gin.Context) {
	var err error
	var result RenameGuitarClassResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req RenameGuitarClassReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = pic.RenameGuitarClass(ctx, req.RenameGuitarClassPara)
	return
}

type (
	RenameGuitarClassReq struct {
		pic.RenameGuitarClassPara
	}
	RenameGuitarClassResp struct {
	}
)
