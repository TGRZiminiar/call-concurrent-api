package google

import (
	"fmt"
	"math/rand"
	"time"
)

type Result3 string
type Search3 func(query string) Result3

var (
	Web3   = fakeSearch3("web3")
	Image3 = fakeSearch3("image3")
	Video3 = fakeSearch3("video3")
)

func fakeSearch3(kind string) Search3 {
	return func(query string) Result3 {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result3(fmt.Sprintf("%s result3 for %q\n", kind, query))
	}
}

// I don't want to wait for slow server
func Google3(query string) []Result3 {
	c := make(chan Result3)

	// each search3 performs in a goroutine
	go func() {
		c <- Web3(query)
	}()
	go func() {
		c <- Image3(query)
	}()
	go func() {
		c <- Video3(query)
	}()

	var result3s []Result3

	// the global timeout for 3 queries
	// it means after 50ms, it ignores the result3 from the server that taking response greater than 50ms
	//
	timeout := time.After(550 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			result3s = append(result3s, r)
		// this line ignore the slow server.
		case <-timeout:
			fmt.Println("timeout")
			return result3s
		}
	}

	return result3s
}

func Run3() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result3s := Google3("golang")
	elapsed := time.Since(start)
	fmt.Println(result3s)
	fmt.Println(elapsed)
}
