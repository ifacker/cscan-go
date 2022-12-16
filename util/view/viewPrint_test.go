package view

import (
	"fmt"
	"testing"
)

func TestPrintlnSuccess(t *testing.T) {
	PrintlnSuccess(fmt.Sprintf("%s, %s", "fuck", "fucks"))
}

func TestPrintlnFailed(t *testing.T) {
	PrintlnFailed("fuck, fuck")
}

func TestPrintlnInfo(t *testing.T) {
	PrintlnInfo("fuck")
}

func TestPrintlnError(t *testing.T) {
	PrintlnError("err")
}
