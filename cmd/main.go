package main

import (
	"buildgoprogram/internal"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/**
为了方便修改一些值定义成常量
 */
const title = "Build Go Program"
const usage = "Automatic directory creation"
const version = "1.0.1"

/**
初始化
 */
func main() {
	app := &cli.App{
		Name:    title,
		Usage:   usage,
		Version: version,
		Action:  buildAction,

		Commands: buildCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/**
具体命令执行的操作
 */
func buildCommand() []*cli.Command {
	return []*cli.Command{
		{
			//命令 名
			Name:    "build",
			//命令 其他触发方式
			Aliases: []string{"b"},
			//命令 介绍
			Usage:   "Automatic directory creation",
			// 命令触发方法
			Action:  internal.BuildProgram,
		},
		{
			Name:    "time",
			Aliases: []string{"t"},
			Usage:   "Print Time",
			Action:  internal.PrintTime,
		},
	}
}

/**
未输出命令执行
 */
func buildAction(c *cli.Context) error {
	fmt.Println("Welcome to use!")
	return nil
}
