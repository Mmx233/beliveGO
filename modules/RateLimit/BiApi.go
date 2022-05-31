package RateLimit

import "time"

var BiApi = make(chan bool)

func init() {
	go func() {
		BiApi <- true
		time.Sleep(time.Second + time.Millisecond*100)
	}()
}
