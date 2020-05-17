package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

/**
打印当前时间 当时写的一个demo
 */
func PrintTime(c *cli.Context) error {
	now := time.Now()
	fmt.Printf("当前时间:%d-%d-%dT %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return nil
}
