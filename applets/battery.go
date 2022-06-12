package applets

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const batteryPath = "/sys/class/power_supply/"

func batteryApplet() (string, error) {
	entries, err := ioutil.ReadDir(batteryPath)
	if err != nil {
		return "", err
	}
	var ret string
	for index, i := range entries {
		if strings.HasPrefix(i.Name(), "BAT") {
			icon := "    "
			status, err := ioutil.ReadFile(filepath.Join(batteryPath,
				i.Name(), "status"))
			if err != nil {
				return "", err
			}
			if string(status) == "Charging\n" {
				icon = "   "
			}
			cap, err := ioutil.ReadFile(filepath.Join(batteryPath,
				i.Name(), "capacity"))
			capS := strings.Trim(string(cap), "\n")
			if index > 0 {
				ret = fmt.Sprintf("%s %s %s%%", ret, icon, capS)
			} else {
				ret = fmt.Sprintf("%s %s%%", icon, capS)
			}
		}
	}
	return ret, nil
}
