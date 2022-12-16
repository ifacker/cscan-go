package file

import (
	"cscan/util/log"
	"io"
	"os"
)

// 读取文件
func ReadFile(filePath string) []byte {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Print(err)
	}
	body, err := io.ReadAll(file)
	if err != nil {
		log.Print(err)
	}
	return body
}
