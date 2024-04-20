package common

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
)

type (
	uniqueKeyExt struct {
		data int64
		lk   sync.RWMutex
	}
)

var uniqueKey uniqueKeyExt

func GetUniqueKey() int64 {
	uniqueKey.lk.Lock()
	defer uniqueKey.lk.Unlock()
	now := time.Now().Unix()
	if now > uniqueKey.data {
		uniqueKey.data = now
	} else {
		uniqueKey.data++
	}
	return uniqueKey.data
}

func GetUniqueKeyStr() string {
	uniqueKey.lk.Lock()
	defer uniqueKey.lk.Unlock()
	now := time.Now().Unix()
	if now > uniqueKey.data {
		uniqueKey.data = now
	} else {
		uniqueKey.data++
	}
	return time.Unix(uniqueKey.data, 0).Format("20060102150405")
}

func GetUniqueMd5() string {
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

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
