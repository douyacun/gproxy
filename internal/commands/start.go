package commands

import (
	"dyc/internal/initialize"
	"github.com/urfave/cli"
)

var Start = cli.Command{
	Name:   "start",
	Usage:  "",
	Action: startAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "env",
			Usage:    "-env [debug, prod]",
			Required: true,
		},
	},
}

func startAction(c *cli.Context) (err error) {
	// 加载配置文件
	initialize.Init(c.String("env"))
	// 启动web服务
	initialize.Server()

	return nil
}
