//go:build linux

package applet

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
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
			if err != nil {
				return "", err
			}
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

const brightnessPath = "/sys/class/backlight/"

func getBrightnessInfo(name string) (brightness int, err error) {
	bp := filepath.Join(brightnessPath, name, "brightness")
	mbp := filepath.Join(brightnessPath, name, "max_brightness")
	bs, err := ioutil.ReadFile(bp)
	if err == nil {
		var mbs []byte
		mbs, err = ioutil.ReadFile(mbp)
		if err == nil {
			brightness, err = strconv.Atoi(strings.Trim(string(bs), "\n"))
			if err == nil {
				var max int
				max, err = strconv.Atoi(strings.Trim(string(mbs), "\n"))
				brightness = int(float32(brightness) / float32(max) * 100)

			}
		}
	}
	return brightness, err

}
func brightnessApplet() (string, error) {
	entries, err := ioutil.ReadDir(brightnessPath)
	if err != nil {
		return "", err
	}
	var ret string
	for index, i := range entries {
		bright, err := getBrightnessInfo(i.Name())
		if err != nil {
			return "", err
		}
		fmt.Println(bright)
		if index > 0 {
			ret = fmt.Sprintf("%s   %d%%", ret, bright)
		} else {
			ret = fmt.Sprintf("  %d%%", bright)
		}
	}
	return ret, nil
}

const audioCmd string = "echo -n \"  \" $(amixer get Master |grep % |sed -e 's/\\].*//' |sed -e 's/.*\\[//')"

var audioApplet = cmdApplet(audioCmd)
