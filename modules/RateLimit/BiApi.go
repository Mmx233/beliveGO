package RateLimit

import "time"

var BiApi = make(chan bool, 8)

func init() {
	go func() {
		for i := 0; i < cap(BiApi); i++ {
			BiApi <- true
		}
		for {
			BiApi <- true
			time.Sleep(time.Second + time.Millisecond*10)
		}
	}()
}
