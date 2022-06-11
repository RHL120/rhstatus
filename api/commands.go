package api

import (
	"fmt"

	"github.com/RHL120/rhstatus/applets"
)

type Command struct {
	function func(arg []string) func(...interface{}) error
	argCount uint8
}

func toggleApplet(arg []string) func(...interface{}) error {
	return func(...interface{}) error {

		applet := applets.FindApplet(arg[0])
		if applet == nil {
			return fmt.Errorf("Could not find applet %s", arg[1])
		}
		applet.Toggle()
		return nil
	}
}

func turnApplet(arg []string) func(...interface{}) error {
	return func(...interface{}) error {
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

func refresh(arg []string) func(...interface{}) error {
	return func(params ...interface{}) error {
		applets.Render()
		return nil
	}
}

var commands map[string]Command = map[string]Command{
	"toggle":  {function: toggleApplet, argCount: 1},
	"turn":    {function: turnApplet, argCount: 2},
	"refresh": {function: refresh, argCount: 0},
}
