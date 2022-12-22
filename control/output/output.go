package output

import (
	"cscan/config"
	"cscan/util/file"
	"fmt"
	"strings"
)

// OutputFile 判断对应的格式文件，并按照对应的格式进行输出
func OutputFile(filepath string, ipOption *config.IpOption, mode int) {
	// 如果没有后缀，自动添加后缀
	if len(filepath) >= 1 && len(filepath) < 5 {
		filepath += ".txt"
	} else if len(filepath) < 1 {
		return
	}
	// 获取后缀
	suffix := filepath[len(filepath)-3:]
	if strings.ToLower(suffix) == "csv" {
		printTmp := ""
		switch mode {
		case config.MODE_PORT:
			if ipOption.Status {
				printTmp += fmt.Sprintf("%s,%d,open\n", ipOption.Ip, ipOption.Port)
			} else {
				//printTmp += fmt.Sprintf("%s,%d,close\n", ipOption.Ip, ipOption.Port)
			}
			file.WriteFile(filepath, []byte(printTmp))
		case config.MODE_WEB:
			printTmp += fmt.Sprintf("%s,%d,%s,%s,%s,%d,%s\n", ipOption.WebInfo.Url, ipOption.WebInfo.Code, ipOption.WebInfo.Title, ipOption.WebInfo.Server, ipOption.WebInfo.XPoweredBy, ipOption.WebInfo.Len, ipOption.WebInfo.Url302Jump)
			file.WriteFile(filepath, []byte(printTmp))
		case config.MODE_CRACK:
			// 待开发...
		}
	} else {
		printTmp := ""
		switch mode {
		case config.MODE_PORT:
			if ipOption.Status {
				printTmp += fmt.Sprintf("%s:%d --> open\n", ipOption.Ip, ipOption.Port)
			} else {
				//printTmp += fmt.Sprintf("%s:%d --> close\n", ipOption.Ip, ipOption.Port)
			}
			file.WriteFile(filepath, []byte(printTmp))
		case config.MODE_WEB:
			printTmp += fmt.Sprintf("%s\t%d\t%s\t%s\t%s\t%d\t%s\n", ipOption.WebInfo.Url, ipOption.WebInfo.Code, ipOption.WebInfo.Title, ipOption.WebInfo.Server, ipOption.WebInfo.XPoweredBy, ipOption.WebInfo.Len, ipOption.WebInfo.Url302Jump)
			file.WriteFile(filepath, []byte(printTmp))
		case config.MODE_CRACK:
			// 待开发...
		}
	}
}
