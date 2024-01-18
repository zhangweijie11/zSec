package main

import (
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/cmd"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "traffic-analysis server"
	app.Author = "zhangweijie"
	app.Version = "20230118"
	app.Usage = "traffic-analysis server"
	app.Commands = []cli.Command{cmd.Start}
	app.Flags = append(app.Flags, cmd.Start.Flags...)
	_ = app.Run(os.Args)
}
