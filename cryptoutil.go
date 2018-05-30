package gocommon

import (
	"crypto/md5"
	"encoding/hex"
)

// ToMd5Hash return md5 hash code of string
func ToMd5Hash(text string) string {
	if len(text) < 0 {
		return text
	}
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
