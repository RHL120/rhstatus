package applets

import (
	"fmt"
	"os/exec"
	"strings"
)

type Applet struct {
	Name           string
	UpdateInterval uint
	Enabled        bool
	Function       func(...interface{}) (string, error)
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

var Applets []Applet = []Applet{
	{Name: "battery", UpdateInterval: 120, Enabled: true, Function: batteryApplet},
	{Name: "date", UpdateInterval: 60, Enabled: true, Function: dateApplet},
	{Name: "time", UpdateInterval: 60, Enabled: true, Function: timeApplet},
}

func PrintApplets() {
	for _, i := range Applets {
		ret, err := i.Function()
		if err == nil {
			fmt.Printf("%s |", ret)
		}
	}
}
