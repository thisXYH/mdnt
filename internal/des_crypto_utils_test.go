package internal

import (
	"crypto/rand"
	"testing"
)

func TestDesCrypto(t *testing.T) {
	src := "锄禾日当午，汗滴禾下土"
	key := make([]byte, 8)
	rand.Read(key)

	iv := make([]byte, 8)
	rand.Read(iv)

	dst := EncyptogDES([]byte(src), key, iv)
	src1 := DecrptogDES(dst, key, iv)
	if src != string(src1) {
		t.Fail()
	}
}
