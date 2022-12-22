package file

import (
	"cscan/config"
	"log"
	"os"
)

// 写入文件
func WriteFile(filepath string, content []byte) bool {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil && config.Debug {
		log.Println(err)
		return false
	}
	file.Write(content)
	return true
}
