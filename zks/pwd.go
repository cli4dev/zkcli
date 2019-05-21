package zks

import (
	"github.com/micro-plat/zkcli/rgsts"
	"github.com/urfave/cli"
)

func init() {
	rgsts.Register(
		cli.Command{
			Name:   "pwd",
			Usage:  "显示当前路径",
			Action: pwdAction,
		})
}

//Action .
func pwdAction(c *cli.Context) (err error) {
	npath := rgsts.GetRoot()
	rgsts.Log.Info(npath)
	return nil
}
