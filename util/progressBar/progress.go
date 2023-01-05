package progressBar

import (
	"github.com/gookit/color"
	"sync"
	"time"
)

// 进度条
func Progress(viewText string, channel chan bool, wg *sync.WaitGroup) {

	// 显示加载界面
	color.C256(226).Print("  " + viewText)
out:
	for {
		select {
		case <-channel:
			break out
		default:
			color.C256(226).Print(".")
			time.Sleep(1 * time.Second)
		}
	}
	//color.C256(226).Println("OK!\n")
	wg.Done()
}
