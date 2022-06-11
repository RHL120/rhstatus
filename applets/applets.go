package applets

import "fmt"

type Applet struct {
	name           string
	UpdateInterval uint
	Enabled        bool
	Function       func(...interface{}) (string, error)
}

var Applets []Applet = []Applet{
	{UpdateInterval: 120, Enabled: true, Function: batteryApplet},
	{UpdateInterval: 60, Enabled: true, Function: dateApplet},
	{UpdateInterval: 60, Enabled: true, Function: timeApplet},
}

func PrintApplets() {
	for _, i := range Applets {
		ret, err := i.Function()
		if err == nil {
			fmt.Printf("%s |", ret)
		}
	}
}
