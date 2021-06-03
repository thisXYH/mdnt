package internal

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

//padding 分组对齐, 填充模式 pkcs7padding
func padding(src []byte, blockSize int) []byte {
	count := blockSize - len(src)%blockSize
	tail := bytes.Repeat([]byte{byte(count)}, count)
	return append(src, tail...)
}

// unPadding 移除填充的字节
func unPadding(src []byte) []byte {
	n := len(src)
	count := int(src[n-1])
	realSrc := src[:n-count]
	return realSrc
}

// EncyptogDES
func EncyptogDES(src, key, iv []byte) []byte {
	block, _ := des.NewCipher(key)
	src1 := padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(src1))
	blockMode.CryptBlocks(dst, src1)
	return dst
}

// DecrptogDES
func DecrptogDES(src, key, iv []byte) []byte {
	block, _ := des.NewCipher(key)
	blockeMode := cipher.NewCBCDecrypter(block, iv)
	// 分组计算，直接把计算结果写到原切片上。
	blockeMode.CryptBlocks(src, src)
	dst := unPadding(src)
	return dst
}

// TryDecrptogDES 不直接Panic
func TryDecrptogDES(src, key, iv []byte) (dst []byte, err error) {
	// 解密失败
	defer func() {
		if p := recover(); p != nil {
			err = p.(error)
		}
	}()

	return DecrptogDES(src, key, iv), nil
}
