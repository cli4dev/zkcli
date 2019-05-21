package zks

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/micro-plat/zkcli/rgsts"
	"github.com/urfave/cli"
)

func init() {
	rgsts.Register(
		cli.Command{
			Name:   "cd",
			Usage:  "进入目录",
			Action: cdAction,
		})
}

//Action .
func cdAction(c *cli.Context) (err error) {
	p, err := getCurrentPath(c)
	if err != nil {
		return err
	}
	if p == "" {
		return
	}
	rgsts.Root = p
	return nil
}
func getCurrentPath(c *cli.Context) (string, error) {
	npath := rgsts.GetRoot()
	curr := c.Args().Get(0)
	switch {
	case curr == "":
		return "", nil
	case strings.HasPrefix(curr, "/"):
		return curr, nil
	case curr == "../" || curr == "..":
		items := strings.Split(rgsts.GetRoot(), "/")
		if len(items) == 1 {
			return "", fmt.Errorf("未指定路径或名称")
		}
		return filepath.Join("/", filepath.Join(items[:len(items)-1]...)), nil
	default:

		paths, _, err := rgsts.Rgst.GetChildren(npath)
		if err != nil {
			return "", err
		}
		for _, p := range paths {
			if strings.Contains(p, curr) {
				npath = filepath.Join(npath, p)
				if b, _ := rgsts.Rgst.Exists(npath); !b {
					rgsts.Log.Printf("%s 不存在", npath)
					return "", fmt.Errorf("%s 不存在", npath)
				}
				return npath, nil
			}
		}
		return "", fmt.Errorf("%s 未找到", curr)
	}
}
