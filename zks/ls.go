package zks

import (
	"fmt"
	"strings"

	"github.com/micro-plat/zkcli/rgsts"
	"github.com/urfave/cli"
)

func init() {
	rgsts.Register(
		cli.Command{
			Name:   "ls",
			Usage:  "列出当前目录所有子节点",
			Action: lsAction,
		})
}

//Action .
func lsAction(c *cli.Context) (err error) {
	npath := rgsts.GetRoot()
	if c.NArg() > 1 {
		npath = fmt.Sprintf("%s/%s", npath, strings.Trim(c.Args().Get(1), "/"))
	}
	paths, _, err := rgsts.Rgst.GetChildren(npath)
	if err != nil {
		return err
	}
	rgsts.Root = npath
	rgsts.Log.Info(getLines(paths))
	return nil
}
func getLines(input []string) string {
	var buff strings.Builder
	max := getMaxLen(input)
	for j, i := range input {
		buff.WriteString(i)
		buff.WriteString(strings.Repeat(" ", max-len(i)+1))
		if j%4 == 0 {
			buff.WriteString("\n")
		}
	}
	return buff.String()
}
func getMaxLen(input []string) int {
	max := 0
	for _, i := range input {
		if len(i) > max {
			max = len(i)
		}
	}
	return max
}
