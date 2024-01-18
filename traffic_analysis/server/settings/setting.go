package settings

import (
	"github.com/zhangweijie11/zSec/traffic_analysis/server/util"
	"gopkg.in/ini.v1"
)

var (
	Cfg    *ini.File
	DEBUG  bool
	SECRET string
)

func init() {
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	if err != nil {
		util.Log.Panicln(err)
	}

	DEBUG = Cfg.Section("").Key("DEBUG").MustBool(true)
	SECRET = Cfg.Section("").Key("SECRET_KEY").MustString("")
}
