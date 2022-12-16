package config

import "time"

// 扫描默认端口
var Ports = []int{80, 81, 82, 83, 84, 85, 86, 87, 89, 88, 443, 8443, 7001, 7080, 7090, 8000, 8008, 8888, 8070, 8080,
	8081, 8082, 8083, 8084, 8085, 8086, 8087, 8088, 8089, 8090, 8161, 9001, 9090, 9443, 21, 22, 445, 1100, 1433, 1434,
	1521, 3306, 3389, 3399, 6379, 8009, 9200, 11211, 27017, 50070}

// 扫描特殊端口
var Ports_other = []int{21, 22, 445, 1100, 1433, 1434, 1521, 3306, 3389, 3399, 6379, 8009, 9200, 11211, 27017, 50070}

// UA
var User_Agent = []string{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36"}

var (
	// 代理
	Proxy = ""

	// 超时时间
	TimeOutSet = 10
	TimeOut    = time.Second * time.Duration(TimeOutSet)

	// Debug 模式开关
	Debug = false

	// 最大线程数
	ThreadMax = 50

	// 展示所有细节
	ViewAll = false

	// 带默认端口扫描
	DefaultPorts string

	// 不带默认端口扫描
	ForbidPorts string

	// IP 文件路径
	IpFilePath string

	// IP 或 IP 段
	IpInfo string
)
