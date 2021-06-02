package encrypto

import (
	"crypto/md5"
	"os"
	"path/filepath"

	"github.com/thisXYH/mdnt/internal"
)

type Options struct {
	FilePath string
	Password string
}

func Execute(p Options) error {
	salt := `~@#.s`
	ext := `.mdnt`
	p1 := salt + p.Password + salt
	// 16 bytes
	p1Md5 := md5.Sum([]byte(p1))
	key, iv := p1Md5[:8], p1Md5[8:16]
	src, err := os.ReadFile(p.FilePath)
	if err != nil {
		return err
	}

	// 加密文件 => 解密
	if filepath.Ext(p.FilePath) == ext {
		// 解密
		plain := internal.DecrptogDES(src, key, iv)
		os.WriteFile(p.FilePath[:len(p.FilePath)-len(ext)], plain, 0666)

		// 删除解密文件
		os.Remove(p.FilePath)
	} else {
		// 加密
		dst := internal.EncyptogDES(src, key, iv)
		os.WriteFile(p.FilePath+ext, dst, 0666)

		// 删除原文件
		os.Remove(p.FilePath)
	}
	return nil
}
