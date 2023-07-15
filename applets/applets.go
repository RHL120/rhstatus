package applets

import (
	"fmt"
	"os/exec"

	"github.com/RHL120/rhstatus/X"
)

//Contains info about the applet
type Applet struct {
	//How will the applet be refrenced by the server. Name shouldn't
	//contain spaces
	Name string
	//should this applet be shown by default
	Enabled bool
	//the function that produces the text to be put on the status bar
	function func() (string, error)
}


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

var applets []*Applet = []*Applet{
	{Name: "caps", Enabled: false, function: constantApplet(" Caps")},
	{Name: "brightness", Enabled: true, function: brightnessApplet},
	{Name: "audio", Enabled: true, function: audioApplet},
	{Name: "battery", Enabled: true, function: batteryApplet},
	{Name: "date", Enabled: true, function: dateApplet},
	{Name: "time", Enabled: true, function: timeApplet},
}

//Toggle applet.Enabled
func (applet *Applet) Toggle() {
	applet.Enabled = !applet.Enabled
}

//run the applets in applets.applets and put their return value onto the status
//bar
func Render() {
	var status string
	for _, i := range applets {
		if i.Enabled {
			ret, err := i.function()
			if err != nil {
				fmt.Printf("Failed to run applet %s because %v\n",
					i.Name, err)
				continue
			}
			if ret != "" {
				status = fmt.Sprintf("%s  │  %s", status, ret)
			}
		}
	}
	X.UpdateStatus(status)
}

//look for the applet with name "name" in applets.applets and return a pointer
//to it
func FindApplet(name string) *Applet {
	for _, i := range applets {
		if i.Name == name {
			return i
		}
	}
	return nil
}
