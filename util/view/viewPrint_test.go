package view

import (
	"fmt"
	"github.com/gookit/color"
	"testing"
)

func TestPrintlnSuccess(t *testing.T) {
	PrintlnSuccess(fmt.Sprintf("%s, %s", color.C256(165).Sprintf("fuck"), color.C256(214).Sprintf("fucks")))
	PrintlnSuccess("a" + color.C256(165).Sprintf("fuck") + "c")
	a := "http://www.baidu.com:80"
	b := color.C256(46).Sprintf("[") + color.C256(165).Sprintf("%d", 200) + color.C256(46).Sprintf("]")
	c := "https://www.baidu.com:443"
	d := "[200]"
	PrintlnSuccess(c + "\t" + d)
	PrintlnSuccess(a + "\t" + b)
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
