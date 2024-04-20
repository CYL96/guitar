package fox4

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"server/src/common"
	"server/src/mod/intf"
)

func (f *Fox4ConfigExt) GetModId() int64 {
	return 1
}

func (f *Fox4ConfigExt) GetModName() string {
	return "fox4"
}

func (f *Fox4ConfigExt) GetModReqURl() string {
	return "https://www.fox4.cn"
}
func (f *Fox4ConfigExt) GetModDownloadPreUrl() string {
	return "https://www.fox4.cn/"
}

func (f *Fox4ConfigExt) GetUrlInfo(url string) (info intf.GuitarGetInfo, err error) {
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
	info.Name = f.readTitle(data)
	if info.Name == "" {
		info.Name = "未知_" + time.Now().Format("20060102150405")
	}

	info.PicList = f.readPng(data)
	return
}

func (f *Fox4ConfigExt) DownLoad(info intf.GuitarGetInfo) (err error) {
	return common.DownloadPic(info.Name, info.PicList)

}

func (f *Fox4ConfigExt) readTitle(data []byte) string {
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

func (f *Fox4ConfigExt) readPng(data []byte) []string {
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
					pngList = append(pngList, f.GetModDownloadPreUrl()+v[start+len(picKey):end])
				}
			}

		} else if strings.Contains(v, "<div class=\"post-content no-indent\">") {
			//     标题
			flag = true
		}

	}

	return pngList
}
