package applets

import "fmt"

type Applet struct {
	Name           string
	UpdateInterval uint
	Enabled        bool
	Function       func(...interface{}) (string, error)
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
