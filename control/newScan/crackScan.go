package newScan

import (
	"cscan/config"
	"cscan/control/crack"
)

// 对一些特殊端口进行暴力破解
func crackStart(ipOption *config.IpOptions) {
	if config.NotCrack {
		return
	}
	crack.StartCrack(ipOption)
}
