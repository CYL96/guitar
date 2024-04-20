/**************************************************
*文件名：code.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/2/28
**************************************************/

package hd

import (
	"fmt"
	"sort"
)

type (
	STATE_CODE int
)

const (
	StatePanic  STATE_CODE = 9999 // 服务器崩溃
	StateOk     STATE_CODE = 0    // 成功
	StateFailed STATE_CODE = 1    // 失败
)

var (
	stateToMessage = map[STATE_CODE]string{
		StatePanic:  "服务器崩溃",
		StateOk:     "成功",
		StateFailed: "失败",
	}
)

// String
// Description: 转换为string
// Author: CYL96
// Date: 2023-08-23 11:06:16
// Receiver c STATE_CODE
// Return string :
func (c STATE_CODE) String() string {
	return stateToMessage[c]
}

// Int
// Description: 转换为int
// Author: CYL96
// Date: 2023-08-23 11:06:16
// Receiver c STATE_CODE
// Return string :
func (c STATE_CODE) Int() int {
	return int(c)
}

// IsErr
// Description: 是否错误
// Author: CYL96
// Date: 2023-08-23 13:48:53
// Receiver c STATE_CODE
// Return bool :
func (c STATE_CODE) IsErr() bool {
	return c != StateOk
}

// GetStateDefine
// Description: 获取状态码对应的信息
// Author: CYL96
// Date: 2023-08-11 11:24:29
// Return map[STATE_CODE]string :
func GetStateDefine() map[STATE_CODE]string {
	return stateToMessage
}

// printStateCode
// Description: 状态码打印
// Author: CYL96
// Date: 2023-08-09 13:44:21
func printStateCode() {
	var (
		slice []STATE_CODE
	)
	slice = make([]STATE_CODE, 0, len(stateToMessage))
	for k := range stateToMessage {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i] < slice[j] {
			return true
		}
		return false
	})
	for _, key := range slice {
		fmt.Println(key, ":", stateToMessage[key])
	}
}
