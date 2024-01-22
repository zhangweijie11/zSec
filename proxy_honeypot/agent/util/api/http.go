package api

import (
	"crypto/md5"
	"fmt"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/settings"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	SECRET string
	APIURL string
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

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("server")
	SECRET = sec.Key("SECRET").MustString("SECRET")
	APIURL = sec.Key("API_URL").MustString("http://127.0.0.1/api/send")
}

func Post(data string) (err error) {
	t := time.Now().Format("2006-01-02 15:04:05")
	hostName, _ := os.Hostname()
	_, err = http.PostForm(APIURL, url.Values{"timestamp": {t}, "secureKey": {MakeSign(t, SECRET)}, "data": {data}, "hostname": {hostName}})
	return err
}
