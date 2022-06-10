package applets

import "fmt"

type Applet struct {
	UpdateInterval uint
	Enabled        bool
	Function       func(...interface{}) (string, error)
}

var Applets map[string]Applet = map[string]Applet{
	"date": {UpdateInterval: 60, Enabled: true, Function: dateApplet},
	"time": {UpdateInterval: 60, Enabled: true, Function: timeApplet},
}

func PrintApplets() {
	for _, i := range Applets {
		ret, err := i.Function()
		if err == nil {
			fmt.Printf("%s |", ret)
		}
	}
}
