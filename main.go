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
	for {
		applets.Render()
		time.Sleep(sleepTime)
	}

}
