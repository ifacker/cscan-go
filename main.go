package main

import (
	"cscan/config"
	"cscan/control/newScan"
	"cscan/flag"
	"fmt"
	"time"
)

var startTime time.Time

func init() {
	// 记录开始时间
	startTime = time.Now()
	fmt.Println("程序开始执行时的时间:", startTime.Format(time.RFC3339))

	// 初始化 log 配置文件
	config.LogConfigInit()
	// 打印 logo
	fmt.Println(config.Logo)
	// 初始化 flag
	flag.Init()
}

func main() {
	//scan.StartScans()
	newScan.NewStartScans()

	// 记录结束时间
	endTime := time.Now()
	fmt.Println("程序执行结束时的时间:", endTime.Format(time.RFC3339))

	// 还可以计算执行时间差
	duration := endTime.Sub(startTime)
	fmt.Printf("程序执行总时间: %v\n", duration)
}
