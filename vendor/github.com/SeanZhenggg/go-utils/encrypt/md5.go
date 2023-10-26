package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5String(input string) string {
	hashByte := md5.Sum([]byte(input))
	return hex.EncodeToString(hashByte[:])
}
