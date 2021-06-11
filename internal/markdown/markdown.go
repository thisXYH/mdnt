package markdown

import (
	"fmt"

	"github.com/thisXYH/mdnt/internal"
)

type Options struct {
	MdDir        string //MdDir 文档目录绝对路径
	DoIdSet      bool   //指定要设置id的文档，或者目录
	DoRelPathFix bool
	NewId        bool //生成新Id
}

const IdLen = 32

var ops Options

func Execute(p Options) error {
	ops = p

	if p.NewId {
		fmt.Printf("> id:%s", internal.GenerateHexString(IdLen))
	}

	if p.DoIdSet {

	}

	if p.DoRelPathFix {

	}

	return nil
}
