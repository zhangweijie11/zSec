package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"path/filepath"
)

// md5 function
func MD5(s string) (m string) {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// create a sign by key & md5
func MakeSign(t string, key string) (sign string) {
	sign = MD5(fmt.Sprintf("%s%s", t, key))
	return sign
}

func SplitWhitePorts(ports []string) map[int][]string {
	result := make(map[int][]string)
	total := len(ports)
	batch := 0
	if total%15 == 0 {
		batch = total / 15
		for i := 0; i < batch; i++ {
			result[i] = ports[i*15 : (i+1)*15]
		}
	} else {
		batch = total / 15
		for i := 0; i < batch; i++ {
			result[i] = ports[i*15 : (i+1)*15]
		}
		result[batch] = ports[batch*15 : total]
	}

	return result
}

func GetCurDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		return "", err
	}
	return dir, err
}
