package flag

import (
	"cscan/config"

	"github.com/projectdiscovery/goflags"
)

func Init() {

	flagSet := goflags.NewFlagSet()

	// 对大小写敏感
	flagSet.CaseSensitive = true

	// 打印logo
	flagSet.SetDescription(config.Logo)

	// 创建组

	flagSet.CreateGroup("config", "常用参数",
		flagSet.StringVarP(&config.IpInfo, "ips", "i", "", "需要扫描的 IP、IP段 或 IP范围，如：192.168.1.1, 192.168.1.1/24, 192.168.1.1-20（仅支持 \",\" 逗号分隔）"),
		flagSet.StringVarP(&config.IpFilePath, "localFile", "l", "", "需要导入的 IP 文件"),
		flagSet.StringVarP(&config.DefaultPorts, "dPort", "dp", "", "default ports 指定端口，然后加上默认端口一块扫描（支持 \",\" 逗号 \" \" 空格 \"；\" 分号分隔）"),
		flagSet.StringVarP(&config.ForbidPorts, "fPorts", "fp", "", "forbid ports 指定端口扫描，禁用默认端口（支持 \",\" 逗号 \" \" 空格 \"；\" 分号分隔）"),
		flagSet.StringVarP(&config.Filter, "filter", "f", "", "filter 过滤保留需要的状态码，并打印输出（支持 \",\" 逗号 \" \" 空格 \"；\" 分号分隔）"),
		flagSet.IntVarP(&config.ThreadMax, "thread", "t", config.ThreadMax, "设置最大线程数"),
		flagSet.IntVarP(&config.TimeOutSet, "timeout", "", config.TimeOutSet, "设置超时时长，单位：秒"),
		flagSet.StringVarP(&config.OutPutType, "output", "o", "", "导出检测的结果，目前支持的格式：txt、csv。如：-o outfile.txt、-o outfile.csv"),
		flagSet.BoolVarP(&config.NotCrack, "notCrack", "nc", config.NotCrack, "not crack 禁用暴力破解功能"),
	)

	flagSet.CreateGroup("debug", "调试参数",
		flagSet.BoolVarP(&config.Debug, "debug", "", config.Debug, "Debug 模式，开启后显示所有的日志信息"),
		flagSet.BoolVarP(&config.ViewAll, "view", "", config.ViewAll, "显示并打印所有细节"),
	)

	flagSet.CreateGroup("proxy", "代理设置",
		flagSet.StringVarP(&config.Proxy, "proxy", "", "", "设置代理，如：socks5://localhost:1080 或 http://localhost:8080"),
	)

	flagSet.Parse()
}
