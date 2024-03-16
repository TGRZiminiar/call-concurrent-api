package main

import (
	"log"
	"os"
	"tgrziminiar/concurrent-api/api"
	"tgrziminiar/concurrent-api/google"
	"tgrziminiar/concurrent-api/pingpong"
	quitsignal "tgrziminiar/concurrent-api/quit-signal"
	ringbuffer "tgrziminiar/concurrent-api/ring-buffer"
	"tgrziminiar/concurrent-api/samplebuffer"
	"tgrziminiar/concurrent-api/timeout"
	workerpool "tgrziminiar/concurrent-api/worker-pool"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Error: .env path is invalid")
	}

	switch os.Args[1] {
	case "api":
		api.Run()

	case "samblebuffer":
		samplebuffer.Run()
	case "timeout":
		timeout.Run()
	case "quit-signal":
		quitsignal.Run()
	case "google-1":
		google.Run1()
	case "google-2":
		google.Run2()
	case "google-3":
		google.Run3()
	case "pingpong":
		pingpong.Run()
	case "ring-buffer":
		ringbuffer.Run()
	case "worker-pool":
		workerpool.Run()

	default:
		log.Fatal("Argument is required write some package name to run")
	}

}
