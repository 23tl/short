package encry

import (
	baseMd5 "crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	md5str := baseMd5.New()
	md5str.Write([]byte(str))
	return hex.EncodeToString(md5str.Sum(nil))
}
