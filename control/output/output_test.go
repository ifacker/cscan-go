package output

import (
	"cscan/config"
	"testing"
)

func TestOutputFile(t *testing.T) {
	OutputFile("abcd.CSV", nil, config.MODE_WEB)
}
