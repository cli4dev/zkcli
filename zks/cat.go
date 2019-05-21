package zks

import (
	"github.com/micro-plat/zkcli/rgsts"
	"github.com/urfave/cli"
)

func init() {
	rgsts.Register(
		cli.Command{
			Name:   "cat",
			Usage:  "显示节点内容",
			Action: catAction,
		})
}

//Action .
func catAction(c *cli.Context) (err error) {
	p, err := getCurrentPath(c)
	if err != nil {
		return nil
	}
	if p == "" {
		p = rgsts.GetRoot()
	}
	buff, _, err := rgsts.Rgst.GetValue(p)
	if err != nil {
		return
	}
	rgsts.Log.Infof("%s:", p)
	rgsts.Log.Info(string(buff))
	return
}
