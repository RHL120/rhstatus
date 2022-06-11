package main

import "C"
import (
	"time"

	"github.com/RHL120/rhstatus/X"
	"github.com/RHL120/rhstatus/applets"
)

const sleepTime = 10 * time.Second

func main() {
	defer X.CloseDisplay()
	serverChan := make(chan func() error)
	timeOutChan := make(chan bool)
	go func(c chan bool) {
		for {
			time.Sleep(sleepTime)
			c <- true
		}
	}(timeOutChan)
	for {
		select {
		case <-timeOutChan:
			applets.Render()
		case f := <-serverChan:
			f()
		}
	}

}
