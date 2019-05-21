package zks

import (
	"github.com/micro-plat/zkcli/rgsts"
	"github.com/urfave/cli"
)

func init() {
	rgsts.Register(
		cli.Command{
			Name:   "rm",
			Usage:  "删除节点",
			Action: rmAction,
		})
}

//Action .
func rmAction(c *cli.Context) (err error) {
	p, err := getCurrentPath(c)
	if err != nil {
		return err
	}
	if p == "" {
		return
	}
	err = rgsts.Rgst.Delete(p)
	if err != nil {
		return err
	}
	rgsts.Log.Infof("remove %s\tok", p)
	return
}
