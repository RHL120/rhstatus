package applets

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/RHL120/rhstatus/X"
)

type Applet struct {
	Name     string
	enabled  bool
	function func(...interface{}) (string, error)
}

func (applet *Applet) ToggleApplet() {
	applet.enabled = !applet.enabled
}

var Applets []Applet = []Applet{
	{Name: "battery", enabled: true, function: batteryApplet},
	{Name: "date", enabled: true, function: dateApplet},
	{Name: "time", enabled: true, function: timeApplet},
}

func cmdApplet(cmd string) func(...interface{}) (string, error) {
	args := strings.Split(cmd, " ")
	return func(i ...interface{}) (string, error) {
		cmd := exec.Command(args[0], args[1:]...)
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(output), nil
	}
}

func Render(dpy interface{}) {
	var status string
	for _, i := range Applets {
		ret, err := i.function()
		fmt.Println(i.Name)
		if err != nil {
			fmt.Printf("Failed to run applet %s because %v\n", i.Name, err)
			continue
		}
		status = fmt.Sprintf("%s  |  %s", status, ret)
		X.UpdateStatus(dpy, status)
	}
}
