package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/zkcli/rgsts"
	"github.com/urfave/cli"
)

//VERSION 版本号
var VERSION = "1.0.0"

//Start 启动应用程序
func Start() {
	// defer time.Sleep(time.Second * 5)
	if len(os.Args) < 2 {
		rgsts.Log.Error("请输入zookeeper服务器地址")
		return
	}
	rgst, err := registry.NewRegistryWithAddress(os.Args[1], rgsts.Log)
	if err != nil {
		rgsts.Log.Error(err)
		return
	}
	rgsts.Rgst = rgst

	for {
		time.Sleep(time.Millisecond * 200)
		var input string
		fmt.Printf("zk:%s >", rgsts.GetRoot())
		scanf(&input)
		if err != nil {
			rgsts.Log.Error(err)
			continue
		}
		inputs := []string{os.Args[0]}
		inputs = append(inputs, strings.Split(input, " ")...)
		app := newCliApp()
		if err := app.Run(inputs); err != nil {
			rgsts.Log.Error(err)
			continue
		}
	}
}

//NewCliApp 创建app
func newCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = VERSION
	app.Usage = "zookeeper框架辅助工具，用于管理zookeeper节点"
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help,h",
		Usage: "查看帮助信息",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version,v",
		Usage: "查看版本信息",
	}
	app.Commands = rgsts.GetCmds()
	return app
}
func scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
