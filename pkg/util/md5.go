package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(val string) string {
	m := md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil))
}
