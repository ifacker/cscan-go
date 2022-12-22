package file

import (
	"cscan/config"
	"io"
	"log"
	"os"
)

// 读取文件
func ReadFile(filePath string) []byte {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil && config.Debug {
		log.Println(err)
	}
	body, err := io.ReadAll(file)
	if err != nil && config.Debug {
		log.Println(err)
	}
	return body
}
