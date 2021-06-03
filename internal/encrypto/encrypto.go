package encrypto

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thisXYH/mdnt/internal"
)

type Options struct {
	FilePath string
	Password string
}

// getKeyAndIv 根据密码获取DES 加密的key 和 iv
// 这个逻辑不能变动，否则导致之前的加密文件解不开。
func getKeyAndIv(password string) ([]byte, []byte) {
	salt := `~@#.s`
	p1 := salt + password + salt
	// 16 bytes
	p1Md5 := md5.Sum([]byte(p1))
	return p1Md5[:8], p1Md5[8:16]
}

func Execute(p Options) (err error) {
	ext := `.mdnt`
	key, iv := getKeyAndIv(p.Password)
	src, err := os.ReadFile(p.FilePath)
	if err != nil {
		return err
	}

	// 加密文件 => 解密
	if filepath.Ext(p.FilePath) == ext {
		// 解密
		plain, err := internal.TryDecrptogDES(src, key, iv)
		if err != nil {
			fmt.Fprintln(os.Stderr, "password error~")
			os.Exit(1) //直接给个提示。
		}

		// 写入明文，删除加密文件
		os.WriteFile(p.FilePath[:len(p.FilePath)-len(ext)], plain, 0666)
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
