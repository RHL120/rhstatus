//go:build freebsd

package applets

/*
#include "c.h"
*/
import (
	"C"
)
import "fmt"

func batteryApplet() (string, error) {
	var batt C.battery_t;
	ret := C.get_battery(&batt)
	if ret == 0 {
		return "", nil
	} else if ret == -1 {
		return "", fmt.Errorf("Failed to get battery info")
	}
	icon := "    "
	if batt.charging {
		icon = "   "
	}
	return fmt.Sprintf("%s%v%%", icon, batt.charge), nil
}

func audioApplet() (string, error) {
	vol := C.get_sound();
	if  vol == -1{
		return "", fmt.Errorf("Failed to get volume info");
	}
	if vol == 0 {
		return "  mute", nil
	}
	return fmt.Sprintf("  %v%%", vol), nil

}

func brightnessApplet() (string, error) {
	bright := C.get_brightness();
	if  bright == -1{
		return "", fmt.Errorf("Failed to get brightness info");
	}
	return fmt.Sprintf("  %d%%", bright), nil
}
