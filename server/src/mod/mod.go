package mod

import (
	"strings"

	"server/src/mod/fox4"
	"server/src/mod/intf"
	"server/src/runCtx"
)

var NowModList map[int64]intf.GuitarModIntf

var ModeResultList []intf.GuitarModT

func GetModByHost(host string) intf.GuitarModIntf {
	for _, v := range NowModList {
		if strings.Contains(v.GetModReqURl(), host) {
			return v
		}
	}
	return nil

}

func AddMod(mod intf.GuitarModIntf) {
	if NowModList == nil {
		NowModList = make(map[int64]intf.GuitarModIntf)
	}
	NowModList[mod.GetModId()] = mod
	ModeResultList = append(ModeResultList, intf.GuitarModT{
		ModId:   mod.GetModId(),
		ModName: mod.GetModName(),
		ModUrl:  mod.GetModReqURl(),
	})
}

func init() {
	AddMod(fox4.Fox4Config)
}

func GetModList(ctx *runCtx.RunCtx, para GetModListPara) []intf.GuitarModT {
	return ModeResultList
}

type (
	GetModListPara struct {
	}
)

func GetModGuitarInfo(ctx *runCtx.RunCtx, para GetModGuitarInfoPara) (info intf.GuitarGetInfo, err error) {

	mod := GetModByHost(para.Host)
	if mod == nil {
		err = ctx.ErrorNew("mod not found")
		return

	}
	info, err = mod.GetUrlInfo(para.Url)
	return
}

type (
	GetModGuitarInfoPara struct {
		Host string `json:"host"`
		Url  string `json:"url" default:"" example:""` //
	}
)

func DownloadModGuitarInfo(ctx *runCtx.RunCtx, para DownloadModGuitarInfoPara) (err error) {
	mod := GetModByHost(para.Host)
	if mod == nil {
		err = ctx.ErrorNew("mod not found")
		return

	}
	err = mod.DownLoad(para.GuitarGetInfo)
	return
}

type (
	DownloadModGuitarInfoPara struct {
		Host string `json:"host"`
		intf.GuitarGetInfo
	}
)
