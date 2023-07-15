//go:build openbsd

package applets

/*
#include <sys/types.h>
#include <machine/apmvar.h>
#include <sys/ioctl.h>
const int apm_ioc_getpower = APM_IOC_GETPOWER;
const u_char apm_ac_on = APM_AC_ON;
typedef struct apm_power_info apm_power_info;
int ioctl2(unsigned int d, unsigned long r, apm_power_info *info) {
	return ioctl(d, r, info);
}
*/
import (
	"C"
)
import (
	"fmt"
	"os"
)

func batteryApplet() (string, error) {
	file, err := os.OpenFile("/dev/apm", os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var status C.apm_power_info
	if C.ioctl2(C.uint(file.Fd()), C.apm_ioc_getpower, &status) < 0 {
		return "", fmt.Errorf("Failed to ioctl")
	}
	icon := "    "
	if status.ac_state == C.apm_ac_on {
		icon = "   "
	}
	return fmt.Sprintf("%s%v%%", icon, status.battery_life), nil
}

func brightnessApplet() (string, error) {
	return "", nil
}

func audioApplet() (string, error) {
	return "", nil
}
