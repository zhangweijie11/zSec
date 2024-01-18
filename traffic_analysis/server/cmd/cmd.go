package cmd

import (
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/web"
)

var Start = cli.Command{
	Name:        "start",
	Usage:       "startup zSec traffic server",
	Description: "startup zSec traffic server",
	Action:      web.RunWeb,
	Flags: []cli.Flag{
		boolFlag("debug, d", "debug mode"),
		stringFlag("server, s", "", "http server address"),
		intFlag("port, p", 1024, "http port"),
	},
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
