package settings

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg *ini.File

	// client config
	InterfaceName string
	ManagerUrl    string
	SecKey        string
	ProxyFlag     bool
)

func init() {
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	if err != nil {
		log.Panicln(err)
	}

	sec := Cfg.Section("client")
	InterfaceName = sec.Key("INTERFACE").MustString("en0")
	ManagerUrl = sec.Key("MANAGER_URL").MustString("")
	SecKey = sec.Key("KEY").MustString("")
	ProxyFlag = sec.Key("PROXY_FLAG").MustBool(true)
}
