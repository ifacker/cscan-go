package config

import (
	"fmt"
	"github.com/gookit/color"
)

var logoUp = color.Red.Sprintf("\n                       ..:::::::::..\n                  ..:::aad88x8888baa:::..\n              .::::d:?88888xxx888?::8b::::.\n            .:::d8888:?888xxxxx??a888888b:::.\n          .:::d8888888a8888xxxaa8888888888b:::.\n         ::::dP::::::::88888x88888::::::::Yb::::\n        ::::dP:::::::::Y888888888P:::::::::Yb::::\n       ::::d8::::x::::::Y8888888P:::::x:::::8b::::\n      .::::88::::::::::::Y88888P::::::::::::88::::.\n      :::::Y8baaaaaaaaaa88P:T:Y88aaaaaaaaaad8P:::::\n      :::::::Y88888888888P::|::Y88888888888P:::::::\n      ::::::::::::::::888:::|:::888::::::::::::::::\n      `:::::::::::::::8888888888888b::::::::::::::'\n       :::::::::::::::88888888888888:::::Cscan::::\n        :::::::::::::d88888888888888::::::NB:::::\n         ::::::::::::88::88::88:::88::::::::::::\n          `::::::::::88::88::88:::88::::::::::'\n            `::::::::88::88::P::::88::::::::'\n              `::::::88::88:::::::88::::::'\n                 ``:::::::::::::::::::''\n                      ``:::::::::''\n")

var logoDown = color.Green.Sprintf("\n    =================   WEB Info Scan  ==================\n    =================  Code by %s ==================\n    =================           %s  ==================\n    +++++++++++++++++++++++++++++++++++++++++++++++++++++\n", author, version)

var Logo = fmt.Sprintf("%s%s", logoUp, logoDown)