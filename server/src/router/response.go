/**************************************************
*文件名：response.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/2/28
**************************************************/

package hd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GinResponse struct {
		State STATE_CODE `json:"state"`
		Msg   string     `json:"msg"`
		Data  any        `json:"data"`
	}
)

// NewGinResponse
// Description: 生成返回结构
// Author: CYL96
// Date: 2023-08-07 16:44:31
// Param state : 状态码
// Param msg : 状态信息
// Param data : 返回数据
// Return *GinResponse :
func NewGinResponse(state STATE_CODE, msg string, data any) *GinResponse {
	resp := &GinResponse{
		State: state,
		Msg:   msg,
		Data:  data,
	}
	if resp.Msg == "" {
		resp.Msg = state.String()
	}
	return resp
}

// NewGinResponseData
// Description: 生成返回结构-自动生成状态信息
// Author: CYL96
// Date: 2023-08-07 16:45:15
// Param state : 状态码
// Param data : 数据
// Return *GinResponse :
func NewGinResponseData(state STATE_CODE, data any) *GinResponse {
	return NewGinResponse(state, "", nil)
}

// NewGinResponseEx
// Description: 生成返回结构-无数据
// Author: CYL96
// Date: 2023-08-07 16:45:44
// Param state : 状态码
// Param msg : 状态信息
// Return *GinResponse :
func NewGinResponseEx(state STATE_CODE, msg string) *GinResponse {
	return NewGinResponse(state, msg, nil)
}

// NewGinResponseState
// Description: 生成返回结构-无数据
// Author: CYL96
// Date: 2023-08-07 16:45:44
// Param state : 状态码
// Return *GinResponse :
func NewGinResponseState(state STATE_CODE) *GinResponse {
	return NewGinResponse(state, "", nil)
}

// GinResponseWithState
// Description: 通用成功返回
// Author: CYL96
// Date: 2023-08-07 16:21:44
// Param c : gin.Context
// Param data : 返回的数据
func GinResponseWithState(c *gin.Context, state STATE_CODE) {
	c.JSON(http.StatusOK, NewGinResponseState(state))
	return
}

// GinResponseWithStateAndMsg
// Description: 通用返回
// Author: CYL96
// Date: 2023-08-07 16:21:44
// Param c : gin.Context
// Param data : 返回的数据
func GinResponseWithStateAndMsg(c *gin.Context, state STATE_CODE, msg string) {
	c.JSON(http.StatusOK, NewGinResponse(state, msg, nil))
	return
}

// GinResponseOk
// Description: 通用成功返回
// Author: CYL96
// Date: 2023-08-07 16:21:44
// Param c : gin.Context
// Param data : 返回的数据
func GinResponseOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, GinResponse{
		State: StateOk,
		Msg:   StateOk.String(),
		Data:  data,
	})
	return
}

// GinResponseCodeDesc
// Description: 通用返回
// Author: CYL96
// Date: 2023-08-07 16:22:24
// Param c : gin.Context
// Param httpCode : http状态码
// Param State : 应答码
// param msg : 应答码描述 为空的时候使用默认描述
// Param data : 数据
func GinResponseCodeDesc(c *gin.Context, httpCode int, State STATE_CODE, msg string, data any) {
	resp := GinResponse{
		State: State,
		Msg:   msg,
		Data:  data,
	}
	if msg == "" {
		resp.Msg = StateOk.String()
	}
	c.JSON(httpCode, resp)
	return
}
