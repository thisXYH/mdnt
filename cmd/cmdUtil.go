package cmd

import (
	"flag"
	"fmt"
	"nt/cmd/types"
	"os"
	"path/filepath"
	"strings"
)

// 被 commandUsage 闭包的error
var commandError error

func commandUsage(flag *flag.FlagSet) func() {
	return func() {
		if commandError != nil {
			fmt.Fprintln(os.Stderr, commandError.Error())
		}
		fmt.Fprintln(os.Stderr, "Command", flag.Name(), ":")
		flag.PrintDefaults()
	}
}

// err != nil 时，打印日志并且退出。
func printErrorAndExit(err error, usage func()) {
	commandError = err
	usage()
	os.Exit(99)
}

func getRefImageType(path string) types.RefImageType {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return types.WebImage
	}

	if filepath.IsAbs(path) {
		return types.AbsImage
	} else {
		return types.RelImage
	}
}
