package main

import (
	"cscan/config"
	"cscan/control/scan"
	"cscan/flag"
	"fmt"
)

func init() {
	// 初始化 log 配置文件
	config.LogConfigInit()
	// 打印 logo
	fmt.Println(config.Logo)
	// 初始化 flag
	flag.Init()
}

func main() {
	scan.StartScans()
}
