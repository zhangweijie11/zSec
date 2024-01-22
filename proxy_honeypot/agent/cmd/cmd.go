package cmd

import (
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/util"
)

var Server = cli.Command{
	Name:        "start",
	Usage:       "start proxy agent",
	Description: "start proxy agent",
	Action:      util.Start,
	Flags: []cli.Flag{
		boolFlag("debug, d", "debug mode"),
		intFlag("port, p", 1080, "proxy port"),
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
