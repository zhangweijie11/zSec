package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"strings"
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

func ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func GetCurDir() (error, string) {
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		return err, ""
	}
	return nil, dir
}

func GetHostNameByIp(ip string) (hostname string, err error) {
	addr, err := net.LookupAddr(ip)
	hostname = strings.Join(addr, ",")
	return hostname, err
}
