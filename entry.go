package main

import (
	"fmt"
	"nt/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		helpAndExit()
	}

	command := os.Args[1]
	switch command {
	case cmd.ImgCommand:
		cmd.EntryImgCommand(os.Args[2:])
	default:
		helpAndExit()
	}
}

func helpAndExit() {
	fmt.Fprintf(os.Stderr, `nt markdown 笔记维护工具
项目地址：https://github.com/thisXYH/NoteTools

Usag: nt command [options]

Commands:
	img	图片管理

Detail: nt <command> -h 查看详情
`)
	os.Exit(0)
}
