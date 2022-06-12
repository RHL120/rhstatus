package applets

import (
	"fmt"
	"os/exec"

	"github.com/RHL120/rhstatus/X"
)

type Applet struct {
	//How will the applet be refrenced by the server. Name shouldn't
	//contain spaces
	Name string
	//should this applet be shown by default
	Enabled bool
	//the function that produces the text to be put on the status bar
	function func() (string, error)
}

const audioCmd string = "echo \"  \" $(amixer get Master |grep % |sed -e 's/\\].*//' |sed -e 's/.*\\[//')"

func cmdApplet(cmd string) func() (string, error) {
	return func() (string, error) {
		cmd := exec.Command("sh", "-c", cmd)
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(output), nil
	}
}

func constantApplet(str string) func() (string, error) {
	return func() (string, error) {
		return str, nil
	}
}

var Applets []*Applet = []*Applet{
	{Name: "caps", Enabled: false, function: constantApplet(" Caps")},
	{Name: "audio", Enabled: true, function: cmdApplet(audioCmd)},
	{Name: "battery", Enabled: true, function: batteryApplet},
	{Name: "date", Enabled: true, function: dateApplet},
	{Name: "time", Enabled: true, function: timeApplet},
}

func (applet *Applet) Toggle() {
	applet.Enabled = !applet.Enabled
}

func Render() {
	var status string
	for _, i := range Applets {
		if i.Enabled {
			ret, err := i.function()
			if err != nil {
				fmt.Printf("Failed to run applet %s because %v\n",
					i.Name, err)
				continue
			}
			if ret != "" {
				status = fmt.Sprintf("%s  |  %s", status, ret)
				X.UpdateStatus(status)
			}
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
