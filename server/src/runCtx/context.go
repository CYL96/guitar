/**************************************************
*文件名：contex.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/1/15
**************************************************/

package runCtx

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	SidKey = "sid"
)

type (

	// RunCtx 运行中RunCtx
	RunCtx struct {
		context.Context
		sid string // sid
	}
)

type (
	uniqueKeyExt struct {
		data int64
		lk   sync.RWMutex
	}
)

var uniqueKey uniqueKeyExt

// GetSid
// Description: 获取唯一md5
// Author: CYL96
// Date: 2023-10-10 11:05:47
// Return string :
func GetSid() string {
	uniqueKey.lk.Lock()
	defer uniqueKey.lk.Unlock()
	now := time.Now().Unix()
	if now > uniqueKey.data {
		uniqueKey.data = now
	} else {
		uniqueKey.data++
	}
	str := strconv.FormatInt(uniqueKey.data, 10)
	return str

}

// DefaultContext
// Description: 默认RunCtx
// Author: CYL96
// Date: {date}
// Param ctx *gin.Context :
// Param logger *hclog.HcLog :
// Return *RunCtx :
func DefaultContext() *RunCtx {
	return &RunCtx{
		Context: context.Background(),
		sid:     GetSid(),
	}
}

// FromContext
// Description: 从gin context中获取RunCtx
// Author: CYL96
// Date: {date}
// Param ctx *gin.Context :
// Param logger *hclog.HcLog :
// Return *RunCtx :
func FromContext(ctx context.Context) *RunCtx {
	return &RunCtx{
		Context: ctx,
		sid:     GetSid(),
	}
}

func WithCancel(ctx *RunCtx) (newCtx *RunCtx, cancelFunc context.CancelFunc) {
	newCtx = &RunCtx{}
	newCtx.Context, cancelFunc = context.WithCancel(ctx.Context)
	return
}

// SetSSID
// Description: 设置ssid
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Param sid string :
func (r *RunCtx) SetSSID(ssid string) {
	r.sid = ssid
}

// GetContext
// Description: 获取context
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Return context.Context :
func (r *RunCtx) GetContext() context.Context {
	return r.Context
}

// GetSSID
// Description: 获取ssid
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Return string :
func (r *RunCtx) GetSSID() string {
	return r.sid
}

// _sidStr
// Description: 获取ssid字符串
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Return string :
func (r *RunCtx) _sidStr() string {
	return fmt.Sprintf(" [sid:%s] ", r.sid)
}

// Info
// Description: Info
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Param v ...any :
func (r *RunCtx) Info(v ...any) {
	newV := make([]any, len(v)+4)
	newV[0] = "\u001B[1;32m" + getTime() + " [INFO] "
	newV[1] = getCaller(1)
	newV[2] = r._sidStr()
	copy(newV[3:len(newV)-1], v)
	newV[len(newV)-1] = "\u001B[0m"
	fmt.Println(newV...)
}

// Error
// Description: Error
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Param v ...any :
func (r *RunCtx) Error(v ...any) {
	newV := make([]any, len(v)+4)
	newV[0] = "\u001B[1;31m" + getTime() + " [ERROR] "
	newV[1] = getCaller(1)
	newV[2] = r._sidStr()
	copy(newV[3:len(newV)-1], v)
	newV[len(newV)-1] = "\u001B[0m"
	fmt.Println(newV...)
}

// Warn
// Description: Warn
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Param v ...any :
func (r *RunCtx) Warn(v ...any) {
	newV := make([]any, len(v)+4)
	newV[0] = "\u001B[1;33m" + getTime() + " [WARN] "
	newV[1] = getCaller(1)
	newV[2] = r._sidStr()
	copy(newV[3:len(newV)-1], v)
	newV[len(newV)-1] = "\u001B[0m"
	fmt.Println(newV...)
}

// ErrorNew
// Description: ErrorNew
// Author: CYL96
// Date: {date}
// Receiver r *RunCtx
// Param v ...any :
func (r *RunCtx) ErrorNew(v ...any) (err error) {
	newV := make([]any, len(v)+4)
	newV[0] = "\u001B[1;31m" + getTime() + " [ERROR] "
	newV[1] = getCaller(1)
	newV[2] = r._sidStr()
	copy(newV[3:], v)
	newV[len(newV)-1] = "\u001B[0m"
	fmt.Println(newV...)
	return errors.New(fmt.Sprint(v...))
}
