package intf

type (
	GuitarModT struct {
		ModId   int64  `json:"mod_id" default:"0" example:"0"` //
		ModName string `json:"mod_name"`
		ModUrl  string `json:"mod_url" default:"" example:""` //
	}

	GuitarGetInfo struct {
		Name    string   `json:"name"`
		PicList []string `json:"picList"`
	}

	GuitarModIntf interface {
		GetModId() int64
		GetModName() string
		GetModReqURl() string
		GetModDownloadPreUrl() string
		GetUrlInfo(url string) (GuitarGetInfo, error)
		DownLoad(info GuitarGetInfo) (err error)
	}
)
