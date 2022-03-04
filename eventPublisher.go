package main

import (
	"time"
)

func publish(vid VideoCallStruct) {
	vid.EventName = "start"
	fireEvent(vid)

	for i := 1; i < 10; i++ {
		startHeartbeats(vid)
	}

}

func startHeartbeats(vid VideoCallStruct) {
	vid.EventName = "heartbeat"
	fireEvent(vid)
	time.Sleep(1 * time.Second)
}

func fireEvent(vid VideoCallStruct) {

	// this is just logging now, but it should send it to an http:// endpoint

	prettyLogger(vid)
}
