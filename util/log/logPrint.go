package log

import "cscan/config"
import log2 "log"

// 打印 log
func Print(log any) {
	if config.Debug {
		log2.Println(log)
	}
}
