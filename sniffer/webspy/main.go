package main

import (
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/sniffer/webspy/cmd"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "webSpy"
	app.Author = "zhangweijie"
	app.Version = "2023/01/17"
	app.Usage = "webSpy, Support local and arp spoof mode"
	app.Commands = []cli.Command{cmd.Start}
	app.Flags = append(app.Flags, cmd.Start.Flags...)
	_ = app.Run(os.Args)
}
