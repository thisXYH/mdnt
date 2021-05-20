package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

var (
	// 移除无用
	removeImage bool
)

func init() {
	flag.BoolVar(&removeImage, "ri", false, "删除没有引用的图片资源")

	if len(os.Args) == 1 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	// 移除没有引用了的图片资源。
	if removeImage {
		removeImageFunc()
	}
}

/*
	1. 获取当前目录下的所有图片文件
	2. 解析出markdown笔记中的图片引用
		eg: ![comment](/../../imageName.png)
	3. 移除没有被引用的图片文件
*/
func removeImageFunc() {
	images := make(map[string]string, 100)
	usedImages := make([]string, 0, 100)

	// 遍历当前目录
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		switch ext := filepath.Ext(path); ext {
		case ".png":
			fallthrough
		case ".jpg":
			images[info.Name()] = path

		case ".md":
			byteBuff, _ := os.ReadFile(path)
			content := string(byteBuff)
			re := regexp.MustCompile(`!\[.*\]\(.+/(.+(?:\.png|\.jpg))\)`)

			for _, v := range re.FindAllStringSubmatch(content, -1) {
				usedImages = append(usedImages, v[1])
			}
		}

		return nil
	})

	fmt.Println("当前有", len(images), "个图片文件!")
	for _, v := range usedImages {
		if _, ok := images[v]; ok {
			delete(images, v)
		}
	}

	fmt.Println("移除了", len(images), "个图片文件!")
	for _, val := range images {
		os.Remove(val)
	}
}
