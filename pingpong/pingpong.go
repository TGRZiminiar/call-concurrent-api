package pingpong

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func player(name string, table chan *Ball) {
	for {
		ball := <-table // player grabs the ball
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball // pass the ball
	}
}

func Run() {
	table := make(chan *Ball)

	// 2 go routine running
	// and pass the hits to each other using pointer
	// and after a second just stop the go routine
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball) // game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over, grab the ball
	fmt.Println("end of the process")
}
