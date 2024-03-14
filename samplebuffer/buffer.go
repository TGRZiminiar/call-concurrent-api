package samplebuffer

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Data string
}

func Run() {
	// Create a buffered channel
	// Buffered channel with a capacity of 5
	messageChannel := make(chan Message, 5)

	// WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Server 1: Sender
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			message := Message{Data: fmt.Sprintf("Message %d", i)}
			messageChannel <- message
			// Simulate some work
			time.Sleep(time.Second)
		}
		// Close the channel when done sending
		close(messageChannel)
	}()

	// Server 2: Receiver
	go func() {
		defer wg.Done()
		for msg := range messageChannel {
			// Process the received message
			fmt.Println("Received:", msg.Data)
		}
	}()

	// Wait for both servers to finish
	wg.Wait()
}
