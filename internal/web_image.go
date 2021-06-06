package internal

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var imageContenType = map[string]string{
	`image/png`:                `.png`,
	`image/jpeg`:               `.jpg`,
	`image/gif`:                `.gif`,
	`application/octet-stream`: `.png`,
}

type stringSlice []string

var imageExt = stringSlice{
	`.png`,
	`.jpg`,
	`.gif`,
}

func (s stringSlice) Contain(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// DownWebImage 下载网络图片，返回图片路径，
// 支持不带文件后缀的路径。
func DownWebImage(url, dir string) (string, error) {
	ext := filepath.Ext(url)
	if ext != "" && !imageExt.Contain(ext) {
		ext = "" // 如果后缀拿到的不是图片类型，那么就从Content-Type中拿。
	}

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if ext == "" {
		contentType, ok := response.Header["Content-Type"]
		if !ok {
			return "", fmt.Errorf("unknow MediaType")
		}

		ext, ok = imageContenType[contentType[0]]
		if !ok {
			return "", fmt.Errorf("not image MediaType :%s", contentType[0])
		}
	}
	// 生成路径。
	imagePath := filepath.Join(dir, GenerateHexString(10)+ext)
	f, _ := os.Create(imagePath)
	defer f.Close()
	io.Copy(f, response.Body)

	return imagePath, nil
}

// GenerateHexString 随机生成指定位数(必须是偶数)的HexString
func GenerateHexString(len int8) string {
	if len <= 1 {
		panic("len must more than 1.")
	}

	buf := make([]byte, len/2)
	rand.Read(buf)

	return fmt.Sprintf("%x", buf)
}
