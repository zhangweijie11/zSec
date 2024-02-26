package main

import (
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/zero_trust/zero_trust_proxy/cmd"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	app := cli.NewApp()
	app.Usage = "zero-trust-proxy"
	app.Version = "0.1"
	app.Author = "zhangweijie"
	app.Commands = []cli.Command{cmd.Server}
	app.Flags = append(app.Flags, cmd.Server.Flags...)
	_ = app.Run(os.Args)
}
