package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"math/rand"
	"time"
)

// GetRandomInt 生成指定范围的随机数
func GetRandomInt(down int, up int) int {
	rand.NewSource(time.Now().UnixNano())
	if up < down {
		panic("最大值不允许小于最小值")
	}
	r := rand.Intn(up-down+1) + down
	return r
}

// GetRandomString 获取指定长度字符串
func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// JsonFromObj 将Obj转为json字符串
func JsonFromObj(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

// ObjFromJson 将json字符串转为Obj
func ObjFromJson(j string) interface{} {
	var v interface{}
	_ = json.Unmarshal([]byte(j), &v)
	return v
}

func MapFromJsonIris(j string) iris.Map {
	var m iris.Map
	_ = json.Unmarshal([]byte(j), &m)
	return m
}

// Md5 生成字符串的md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GetFullUrl 获取完整Url
func GetFullUrl(url string) string {
	return IndexUrl + url
}
