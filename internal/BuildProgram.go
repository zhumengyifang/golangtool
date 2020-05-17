package internal

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
)

/**
方便进行其他build选项
 */
func BuildProgram(c *cli.Context) error {
	parameter := c.Args().First()
	switch parameter {
	case "default":
		return defaultBuild(c)
	default:
		return defaultBuild(c)
	}
	return nil
}

/**
基础build 只创建 api build cmd deployments docs examples  internal scripts tools web 文件夹
每个目录的介绍: https://blog.csdn.net/weixin_40165163/article/details/103255972
 */
func defaultBuild(c *cli.Context) error {
	err := pwd()
	if err != nil {
		return err
	}

	cmdStr := "mkdir api build cmd deployments docs examples  internal scripts tools web"
	err = mkdir(cmdStr)
	if err != nil {
		return err
	}

	fmt.Println("successful")
	return nil
}

/**
输出当前所在目录
 */
func pwd() error {
	cmd := exec.Command("sh", "-c", "pwd")
	opBytes, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(opBytes))
	return nil
}

/**
创建所有文件夹
 */
func mkdir(cmdStr string) error {
	if cmdStr == "" {
		return nil
	}

	fmt.Println(cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	var stderr bytes.Buffer
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}
