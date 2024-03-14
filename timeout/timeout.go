package timeout

import (
	"fmt"
	"math/rand"
	"time"
)

func work(text string) <-chan string {
	c := make(chan string)

	go func() {
		i := 0
		for {
			c <- fmt.Sprintf("%s -> %d", text, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			i++
		}
	}()

	return c
}

// let the function to the work with a specific time
// after the limitation of time we end process instanly
// don't care that work done or not
func Run() {

	c := work("testing")
	// set the limitation of time only 3 second
	// after 3 second end the work
	timeout := time.After(3 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println("response -> ", s)
		case <-timeout:
			fmt.Println("time is up!")
			return
		}
	}

}
