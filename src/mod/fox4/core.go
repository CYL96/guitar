/**************************************************
*Copyright(C).2016-2024,瀚辰光翼⽣物科技有限公司
*文件名：core.go
*内容简述：*
*文件历史：
author 李承益 创建 2024/2/19
**************************************************/

package fox4

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

		if flag1 {
			if !strings.Contains(v, "<a class=\"tag\"") {
				break
			}
			start := strings.Index(v, ">")
			if start <= 0 {
				continue
			}
			end := strings.Index(v, "</a>")
			if end <= 0 {
				continue
			}
			if len(titles) > 0 {
				titles += "-"
			}
			titles = strings.ReplaceAll(v[start+1:end], " ", "")

		} else if strings.Contains(v, "<a class=\"tag\"") {
			flag1 = true
		}

	}
	return titles
}

func readPng(data []byte) []string {
	buf := bufio.NewReader(bytes.NewReader(data))
	var flag = false
	pngList := []string{}
	picKey := "<img class=\"post-tab\" src=\""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if flag {

			if strings.Contains(v, "</div>") {
				break
			}

			start := strings.Index(v, picKey)
			if start != -1 {
				end := strings.Index(v, "\" alt=\"")
				if end != -1 {
					pngList = append(pngList, homePage+v[start+len(picKey):end])
				}
			}

		} else if strings.Contains(v, "<div class=\"post-content no-indent\">") {
			//     标题
			flag = true
		}

	}

	return pngList
}
