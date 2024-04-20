package pic

import (
	"os"
	"strings"

	"server/src/common"
	"server/src/runCtx"
)

type (
	GuitarClassT struct {
		Id        int64    `json:"id"`
		ClassName string   `json:"class_name" default:"" example:""` //
		PicList   []string `json:"picList"`
	}
)

func GetGuitarClassList(ctx *runCtx.RunCtx, para GetGuitarClassListPara) (list []GuitarClassT, err error) {
	dir, err := os.ReadDir("Guitar")
	if err != nil {
		ctx.Error(err)
		return
	}
	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}
		if para.Name != "" && para.Name != entry.Name() {
			continue
		}
		if para.Search != "" && !strings.Contains(entry.Name(), para.Search) {
			continue
		}

		picInfo, err := os.ReadDir("Guitar/" + entry.Name())
		if err != nil {
			ctx.Warn(err)
			continue
		}
		var itemInfo GuitarClassT
		itemInfo.Id = common.GetUniqueKey()
		itemInfo.ClassName = entry.Name()
		for _, pic := range picInfo {
			if pic.IsDir() {
				continue
			}
			itemInfo.PicList = append(itemInfo.PicList, "Guitar/"+entry.Name()+"/"+pic.Name())
		}
		list = append(list, itemInfo)

		if para.Name != "" {
			return list, nil
		}
	}
	return
}

type (
	GetGuitarClassListPara struct {
		Name   string `json:"name" default:"" example:""` //
		Search string `json:"search"`
	}
)

func DeleteGuitarClass(ctx *runCtx.RunCtx, para DeleteGuitarClassPara) (err error) {
	if len(para.Name) == 0 {
		return ctx.ErrorNew("参数错误")

	}
	err = os.RemoveAll("Guitar/" + para.Name)
	return
}

type (
	DeleteGuitarClassPara struct {
		Name string `json:"name" default:"" example:""` //
	}
)

func RenameGuitarClass(ctx *runCtx.RunCtx, para RenameGuitarClassPara) (err error) {
	if len(para.Name) == 0 {
		return ctx.ErrorNew("参数错误")

	}
	err = os.Rename("Guitar/"+para.OldName, "Guitar/"+para.Name)
	return
}

type (
	RenameGuitarClassPara struct {
		OldName string `json:"oldName"`
		Name    string `json:"name" default:"" example:""` //
	}
)
