package config

import "log"

// log 格式化配置
func LogConfigInit() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
