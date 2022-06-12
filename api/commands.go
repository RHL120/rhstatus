package api

import (
	"fmt"
	"os"

	"github.com/RHL120/rhstatus/X"
	"github.com/RHL120/rhstatus/applets"
)

type command struct {
	//takes the the arguments and prdouces the function to be called on exec
	//function does not need to check for the argument count this is done
	//by the server
	function func(arg []string) func() error
	//the amount of arguments expected by function
	argCount uint8
}

func toggleApplet(arg []string) func() error {
	return func() error {

		applet := applets.FindApplet(arg[0])
		if applet == nil {
			return fmt.Errorf("Could not find applet %s", arg[1])
		}
		applet.Toggle()
		return nil
	}
}

func turnApplet(arg []string) func() error {
	return func() error {
		var enabled bool
		switch arg[0] {
		case "on":
			enabled = false
		case "off":
			enabled = true
		default:
			return fmt.Errorf("Unknown option %s expected <on> or <off>", arg[0])
		}
		applet := applets.FindApplet(arg[1])
		if applet == nil {
			return fmt.Errorf("Could not find applet %s", arg[1])
		}
		applet.Enabled = enabled
		return nil
	}
}

func refresh(arg []string) func() error {
	return func() error {
		//refresh does nothing since after calling the function
		//main calls render
		return nil
	}
}

func shutdown(arg []string) func() error {
	return func() error {
		X.CloseDisplay()
		os.Exit(0)
		return nil
	}
}

var commands map[string]command = map[string]command{
	"shutdown": {function: shutdown, argCount: 0},
	"toggle":   {function: toggleApplet, argCount: 1},
	"turn":     {function: turnApplet, argCount: 2},
	"refresh":  {function: refresh, argCount: 0},
}
