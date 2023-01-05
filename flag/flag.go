package flag

import (
	"cscan/config"
	"flag"
)

func Init() {

	flag.StringVar(&config.Proxy, "proxy", "", "设置代理，如：socks5://localhost:1080 或 http://localhost:8080")
	flag.IntVar(&config.TimeOutSet, "timeout", config.TimeOutSet, "设置超时时长，单位：秒")
	flag.BoolVar(&config.Debug, "debug", config.Debug, "Debug 模式，开启后显示所有的日志信息")
	flag.IntVar(&config.ThreadMax, "t", config.ThreadMax, "设置最大线程数")
	flag.BoolVar(&config.ViewAll, "v", config.ViewAll, "显示并打印所有细节")
	flag.StringVar(&config.DefaultPorts, "dp", "", "default ports 指定端口，然后加上默认端口一块扫描（支持 \",\" 逗号 \" \" 空格 \"；\" 分号分隔）")
	flag.StringVar(&config.ForbidPorts, "fp", "", "forbid ports 指定端口扫描，禁用默认端口（支持 \",\" 逗号 \" \" 空格 \"；\" 分号分隔）")
	flag.StringVar(&config.IpFilePath, "l", "", "需要导入的 IP 文件")
	flag.StringVar(&config.IpInfo, "i", "", "需要扫描的 IP、IP段 或 IP范围，如：192.168.1.1, 192.168.1.1/24, 192.168.1.1-20（仅支持 \",\" 逗号分隔）")
	flag.StringVar(&config.Filter, "f", "", "filter 过滤保留需要的状态码，并打印输出（支持 \",\" 逗号 \" \" 空格 \"；\" 分号分隔）")
	flag.StringVar(&config.OutPutType, "o", "", "导出检测的结果，目前支持的格式：txt、csv。如：-o outfile.txt、-o outfile.csv")
	flag.BoolVar(&config.NotCrack, "nc", config.NotCrack, "not crack 禁用暴力破解功能")

	flag.Parse()
}
