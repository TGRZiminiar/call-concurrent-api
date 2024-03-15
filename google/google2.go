package google

import (
	"fmt"
	"math/rand"
	"time"
)

type Result2 string
type Search2 func(query string) Result2

var (
	Web2   = fakeSearch2("web2")
	Image2 = fakeSearch2("image2")
	Video2 = fakeSearch2("video2")
)

func fakeSearch2(kind string) Search2 {
	return func(query string) Result2 {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result2(fmt.Sprintf("%s result2 for %q\n", kind, query))
	}
}

func Google2(query string) []Result2 {
	c := make(chan Result2)

	// each search2 performs in a goroutine
	go func() {
		c <- Web2(query)
	}()
	go func() {
		c <- Image2(query)
	}()
	go func() {
		c <- Video2(query)
	}()

	var result2s []Result2
	for i := 0; i < 3; i++ {
		result2s = append(result2s, <-c)
	}

	return result2s
}

func Run2() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result2s := Google2("golang")
	elapsed := time.Since(start)
	fmt.Println(result2s)
	fmt.Println(elapsed)
}
