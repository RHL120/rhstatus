package main

import "C"
import (
	"time"

	"github.com/RHL120/rhstatus/X"
	"github.com/RHL120/rhstatus/applets"
)

const sleepTime = 10 * time.Second

func main() {
	dpy := X.OpenDisplay()
	defer X.CloseDisplay(dpy)
	for {
		applets.Render(dpy)
		time.Sleep(sleepTime)
	}

}
