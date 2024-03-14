package quitsignal

import (
	"fmt"
	"math/rand"
	"time"
)

func work(quit chan struct{}) <-chan string {

	res := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			// case of sending message out of a routine into a main function
			case res <- fmt.Sprintf("%s : %d", "Mix", i):
			// case of when receiving a quit signal
			case <-quit:
				quit <- struct{}{}
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return res
}

// it a signal to do the work until it receive some quit signal
// after it recieve the signal it will terminate all the goroutine in the fn
// :: use case like we want it to do untill someting is done
// :: case of want to do someting with the data that already process instanly not care about the order
func Run() {
	// start of creating a quit chan with a typeof struct
	// using struct because struct use only 0 byte memory
	quit := make(chan struct{})

	res := work(quit)

	// assuming that need to finished this process first
	// and after that terminate the fn work
	for i := 0; i < 3; i++ {
		fmt.Printf("res -> %s\n", <-res)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

	// send a signal to a fn work to terminate the routine
	quit <- struct{}{}

	fmt.Println("quit chan say ->", <-quit)

}
