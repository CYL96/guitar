package cwguitar

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"server/src/common"
)

const (
	PATH = "Guitar"
)

func downLoad(downName, url string) (err error) {
	log.Println("下载：", url)
	var (
		post *http.Request
	)

	post, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	post.Header.Add("Cookie", "PHPSESSID="+CWGuitarConfig.Token+";")
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
			if strings.Contains(v, "class=\"") {
				continue
			}
			titles += v
			if strings.Contains(v, "</h2>") {
				break
			}
		}
		if strings.Contains(v, "<h2 class=\"pt_10 text-ellipsis\">") {
			flag1 = true
		}

	}
	titles = strings.Replace(titles, "</h2>", "", -1)
	titles = strings.Replace(titles, " ", "", -1)
	return titles
}

func readPng(data []byte) []string {
	buf := bufio.NewReader(bytes.NewReader(data))
	var flag = false
	var flag2 = false
	pngStr := ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if flag2 {
			pngStr += v
			if strings.Contains(v, "</div>") {
				break
			}
		}
		if flag && strings.Contains(v, "<div class=\" thumbnail mb_0\">") {
			flag2 = true

		}

		if strings.Contains(v, "<div class=\"clearfix mt_10\">") {
			//     标题
			flag = true
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
