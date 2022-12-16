package file

import (
	"fmt"
	"testing"
)

func TestWriteFile(t *testing.T) {
	body := []byte("hello world")
	b := WriteFile("test.txt", body)
	fmt.Println(b)
}
