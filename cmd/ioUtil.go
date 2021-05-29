package cmd

import (
	"bufio"
	"fmt"
	"os"
)

// 判断文件或者目录是否存在
func isFileOrDirExist(addr string) bool {
	_, err := os.Stat(addr)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func writeToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "file write failed. err: "+err.Error())
	} else {
		defer f.Close()

		bf := bufio.NewWriter(f)
		bf.WriteString(content)
		bf.Flush()
	}
	return err
}
