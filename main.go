package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// 移除无用
	removeImage string
)

func init() {
	flag.StringVar(&removeImage, "ri", ".resources", "图片资源的目录,\n删除没有引用的图片资源")

	if len(os.Args) == 1 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
}
