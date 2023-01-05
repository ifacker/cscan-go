package progressBar

import (
	"sync"
	"testing"
	"time"
)

func TestProgress(t *testing.T) {
	c := make(chan bool)
	wg := &sync.WaitGroup{}
	viewText := "Test"
	wg.Add(1)
	go Progress(viewText, c, wg)

	//go func(c chan bool) {
	//	color.C256(226).Print("  Loading")
	//out:
	//	for {
	//		select {
	//		case <-c:
	//			break out
	//		default:
	//			color.C256(226).Print(".")
	//			time.Sleep(1 * time.Second)
	//		}
	//	}
	//	color.C256(226).Println("OK!\n")
	//	wg.Done()
	//}(c)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
	}

	c <- true

	wg.Wait()
}
