package util

import (
	"crypto/md5"
	"fmt"
)

type Message struct {
	Status  int
	Message string
}

func MakeMd5(srcStr string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(srcStr)))
}

func EncryptPass(src string) string {
	return fmt.Sprintf("%s", MakeMd5(MakeMd5(src)[5:10]))
}
