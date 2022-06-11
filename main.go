package main

import (
	"time"

	"github.com/RHL120/rhstatus/X"
	"github.com/RHL120/rhstatus/api"
	"github.com/RHL120/rhstatus/applets"
)

const sleepTime = 10 * time.Second

func main() {
	defer X.CloseDisplay()
	serverChan := make(chan func() error)
	timeOutChan := make(chan bool)
	go func(c chan bool) {
		for {
			c <- true
			time.Sleep(sleepTime)
		}
	}(timeOutChan)
	go api.RunServer(serverChan)
	for {
		select {
		case <-timeOutChan:
			applets.Render()
		case f := <-serverChan:
			f()
		}
	}

}
