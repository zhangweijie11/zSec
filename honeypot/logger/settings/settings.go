package settings

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	Cfg       *ini.File
	HttpPort  int
	DebugMode bool
	KEY       string
)

func init() {
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)
	if err != nil {
		log.Panic(err)
	}

	sec := Cfg.Section("server")
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	DebugMode = sec.Key("DEBUG_MODE").MustBool(true)
	KEY = sec.Key("KEY").MustString("zSec")
}
