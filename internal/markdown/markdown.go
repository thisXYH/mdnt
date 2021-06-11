package markdown

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/thisXYH/mdnt/internal"
)

type Options struct {
	MdDir        string //MdDir 文档目录绝对路径
	DoIdSet      bool   //给未设置id的文档设置id
	DoRelPathFix bool   //维护文档引用
	NewId        bool   //生成新Id
}

const IdLen = 32

var ops Options

func Execute(p Options) error {
	ops = p

	if p.NewId {
		fmt.Println(newIdWitFormat())
	}

	if p.DoIdSet {
		idSet()
	}

	if p.DoRelPathFix {
		relPathFix()
	}

	return nil
}

func idSet() {
	wg := &sync.WaitGroup{}

	filepath.WalkDir(ops.MdDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || filepath.Ext(d.Name()) != ".md" {
			return nil
		}
		wg.Add(1)
		go idSetCore(path, d.Name(), wg)
		return nil
	})

	wg.Wait()
}

//idSetCore
func idSetCore(path, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	id, moreOneLine := getId(path)
	if id == "" && moreOneLine {
		setId(path)
	}
}

func setId(path string) {
	srcR, _ := os.Open(path)
	srcBr := bufio.NewReader(srcR)
	tempPath := path + ".temp"
	tempW, _ := os.Create(tempPath)

	// copy first line.
	firstLine, _ := srcBr.ReadString('\n')
	tempW.WriteString(firstLine)

	// write id at second line.
	tempW.WriteString(newIdWitFormat() + "\r\n")

	// copy other lines.
	io.Copy(tempW, srcBr)
	srcR.Close()
	tempW.Close()
	os.Rename(tempPath, path)
}

// getId 获取id，并且标记是否文件大于 1 行。
func getId(path string) (string, bool) {
	r, _ := os.Open(path)
	defer r.Close()
	rb := bufio.NewReader(r)

	// id在第二行
	rb.ReadLine()
	buf, _, err := rb.ReadLine()
	if err != nil {
		return "", false
	}

	reg := regexp.MustCompile(`^> +id:([0-9a-f]{32})`)
	result := reg.FindAllStringSubmatch(string(buf), -1)
	if result == nil {
		return "", true
	}

	return result[0][1], true
}

func relPathFix() {
	// TODO
}

func newIdWitFormat() string {
	return fmt.Sprintf("> id:%s", internal.GenerateHexString(IdLen))
}
