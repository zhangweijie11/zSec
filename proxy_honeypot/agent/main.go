package main

import (
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/cmd"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/util"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/vars"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	vars.CurrentDir = util.GetCurDir()
	vars.CaKey = filepath.Join(vars.CurrentDir, "./certs/ca.key")
	vars.CaCert = filepath.Join(vars.CurrentDir, "./certs/ca.cert")

	// log.Logger.Infof("dir: %v, caKey: %v, caCert: %v", vars.CurrentDir, vars.CaKey, vars.CaCert)
}

func main() {
	app := cli.NewApp()
	app.Usage = "proxy agent"
	app.Version = "0.1"
	app.Author = "zhangweijie"
	app.Commands = []cli.Command{cmd.Server}
	app.Flags = append(app.Flags, cmd.Server.Flags...)
	_ = app.Run(os.Args)
}
