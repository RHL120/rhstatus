package applets

import (
	"fmt"
	"os/exec"

	"github.com/RHL120/rhstatus/X"
)

type Applet struct {
	Name     string
	Enabled  bool
	function func(...interface{}) (string, error)
}

const audioCmd string = "echo \"  \" $(amixer get Master |grep % |sed -e 's/\\].*//' |sed -e 's/.*\\[//')"

func cmdApplet(cmd string) func(...interface{}) (string, error) {
	return func(i ...interface{}) (string, error) {
		cmd := exec.Command("sh", "-c", cmd)
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(output), nil
	}
}

var Applets []*Applet = []*Applet{
	{Name: "audio", Enabled: true, function: cmdApplet(audioCmd)},
	{Name: "battery", Enabled: true, function: batteryApplet},
	{Name: "date", Enabled: true, function: dateApplet},
	{Name: "time", Enabled: true, function: timeApplet},
}

func (applet *Applet) Toggle() {
	applet.Enabled = !applet.Enabled
	fmt.Println(*applet)
}

func Render() {
	var status string
	for _, i := range Applets {
		if i.Enabled {
			fmt.Println(i.Name, " ", i.Enabled)
			ret, err := i.function()
			if err != nil {
				fmt.Printf("Failed to run applet %s because %v\n",
					i.Name, err)
				continue
			}
			status = fmt.Sprintf("%s  |  %s", status, ret)
			X.UpdateStatus(status)
		}
	}
}
func FindApplet(name string) *Applet {
	for _, i := range Applets {
		if i.Name == name {
			return i
		}
	}
	return nil
}
