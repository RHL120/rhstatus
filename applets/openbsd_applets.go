//go:build openbsd

package applets

/*
#cgo LDFLAGS: -lsndio
#include <string.h>
#include <stdbool.h>
#include <sys/types.h>
#include <machine/apmvar.h>
#include <sys/ioctl.h>
#include <sndio.h>
typedef struct apm_power_info apm_power_info;
typedef struct sioctl_hdl sioctl_hdl;
const int apm_ioc_getpower = APM_IOC_GETPOWER;
const u_char apm_ac_on = APM_AC_ON;
const char *sio_devany = SIO_DEVANY;
const unsigned int sioctl_read = SIOCTL_READ;
typedef struct apm_power_info apm_power_info;
int ioctl2(unsigned int d, unsigned long r, apm_power_info *info) {
	return ioctl(d, r, info);
}
typedef struct {
	bool mute;
	double volume;
} volume_info;
typedef struct sioctl_hdl sioctl_hdl;
void get_volume_info_cb(void *arg, struct sioctl_desc *desc, int val) {
	volume_info *ret = (volume_info *) arg;
	if (desc && !strcmp(desc->node0.name, "output")) {
		if (!strcmp(desc->func, "level"))
			ret->volume = (double) val / desc->maxval * 100;
		else if (!strcmp(desc->func, "mute"))
			ret->mute = val;

	}
}

volume_info get_volume_info(sioctl_hdl *hdl) {
	volume_info ret = {0};
	sioctl_ondesc(hdl, get_volume_info_cb, (void *)&ret);
	return ret;
}

*/
import (
	"C"
)
import (
	"fmt"
	"math"
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

func audioApplet() (string, error) {
	hdl := C.sioctl_open(C.sio_devany, C.sioctl_read, 1)
	if hdl == nil {
		return "", fmt.Errorf("Failed to open hdl")
	}
	defer C.sioctl_close(hdl)
	vol := C.get_volume_info(hdl)
	if vol.mute {
		return "  mute", nil
	}
	return fmt.Sprintf("  %v%%", int(math.RoundToEven(float64(vol.volume)))), nil
}

const brightnessCmd string = "echo -n \"  \" $(xbacklight |cut  -d . -f 1) %"

var brightnessApplet = cmdApplet(brightnessCmd)
