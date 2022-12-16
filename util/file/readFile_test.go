package file

import (
	"cscan/config"
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	config.Debug = true
	body := ReadFile("../../go.mod")
	fmt.Println(string(body))
}
