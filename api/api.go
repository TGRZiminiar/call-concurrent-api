package api

import (
	"errors"
	"log"
	"time"
)

type (
	Res1 struct {
		Name string
	}

	Res2 struct {
		Age int64
	}
)

func Run() {
	now := time.Now()
	chanRes1 := make(chan *Res1, 1)
	chanErr1 := make(chan error, 1)
	chanRes2 := make(chan *Res2, 1)
	chanErr2 := make(chan error, 1)

	go api1(chanRes1, chanErr1)
	go api2(chanRes2, chanErr2)

	res1 := <-chanRes1
	err1 := <-chanErr1

	res2 := <-chanRes2
	err2 := <-chanErr2

	log.Println("response 1 ->", res1)
	log.Println("error response 1 ->", err1)
	log.Println("\n")
	log.Println("response 2 ->", res2)
	log.Println("error response 2 ->", err2)

	since := time.Since(now)
	log.Println("time use ", since)
}

func api1(chanRes1 chan<- *Res1, chanErr1 chan<- error) {
	log.Println("forcing to delay 1 second")
	time.Sleep(1 * time.Second)
	chanRes1 <- &Res1{
		Name: "mix",
	}
	// chanErr1 <- nil
	chanErr1 <- errors.New("test error msg")
}

func api2(chanRes2 chan<- *Res2, chanErr2 chan<- error) {
	log.Println("forcing to delay 3 second")
	time.Sleep(3 * time.Second)
	chanRes2 <- &Res2{
		Age: 10,
	}
	// chanErr2 <- nil
	chanErr2 <- errors.New("test error msg")

}
