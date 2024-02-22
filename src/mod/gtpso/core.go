/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：core.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package gtpso

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"tool/guitar/src/common"
	"tool/guitar/src/win"
)

func downLoad(open bool, downName, url string) (err error) {
	defer func() {
		if err != nil {
			win.Alter("Download Failed", err.Error())
			return
		}

		if open {
			common.OpenRelativeExplorer(downName)
		}
	}()
	log.Println("下载：", url)
	var (
		post *http.Request
	)

	post, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	post.Header.Add("Cookie", "PHPSESSID="+GtpsoConfig.Token+";")
	res, err := http.DefaultClient.Do(post)
	if err != nil {
		return
	}
	defer res.Body.Close()
	var data []byte
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	if downName == "" {
		downName = readTitle(data)
		if downName == "" {
			downName = time.Now().Format("unk_2006_01_02_15_04_05")
		}
	}

	pngStr := readPng(data)
	if len(pngStr) == 0 {
		err = errors.New("not found guitar sheet")
		return
	}
	err = common.DownloadPic(downName, pngStr)
	if err != nil {
		return
	}
	return
}

func readTitle(data []byte) string {
	buf := bufio.NewReader(bytes.NewReader(data))
	var flag1 = false
	titles := ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)

		if strings.Contains(v, "<title>") {
			flag1 = true
		}
		if flag1 {
			titles += v
			if strings.Contains(v, "</title>") {
				break
			}
		}

	}
	titles = strings.Replace(titles, "<title>", "", -1)
	titles = strings.Replace(titles, "</title>", "", -1)
	titles = strings.Replace(titles, " ", "", -1)
	titles = strings.Replace(titles, "\t", "", -1)
	titles = strings.Replace(titles, "\n", "", -1)
	return titles
}

func readPng(data []byte) []string {
	buf := bufio.NewReader(bytes.NewReader(data))
	var flag = false
	pngStr := ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if strings.Contains(v, "<img") {
			//     标题
			flag = true
		}
		if flag {
			pngStr += v
			if strings.Contains(v, "</div>") {
				break
			}
		}
	}
	pngStr = strings.Replace(pngStr, "</div>", "", -1)
	tmps := strings.Split(pngStr, "\"")
	pngList := []string{}
	for _, v := range tmps {
		if len(v) > 4 && v[:4] == "http" {
			pngList = append(pngList, v)
		}
	}
	return pngList
}
