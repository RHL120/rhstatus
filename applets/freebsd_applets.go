//go:build freebsd

package applets

/*
#include <sys/ioctl.h>
#include <machine/apm_bios.h>

typedef struct apm_info apm_info;
typedef struct apm_pwstatus apm_pwstatus;

const u_int apm_batt_not_present = APM_BATT_NOT_PRESENT;
const unsigned int apmio_getinfo = APMIO_GETINFO;
const u_int apm_batt_charging = APM_BATT_CHARGING;
const u_int apmio_getpwstatus = APMIO_GETPWSTATUS;
const u_int pmdv_batt0 = PMDV_BATT0;

int ioctl2(unsigned int d, unsigned long r, apm_info *info) {
	return ioctl(d, r, info);
}

int ioctl3(unsigned int d, unsigned long r, apm_pwstatus *info) {
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
	var info C.apm_info
	var aps C.apm_pwstatus
	file, err := os.OpenFile("/dev/apm", os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	fd := C.uint(file.Fd());
	defer file.Close()
	if C.ioctl2(fd, C.apmio_getinfo, &info) < 0 {
		return "", fmt.Errorf("Failed to ioctl")
	}
	for i := 0; i < int(info.ai_batteries); i++ {
		aps.ap_device = C.pmdv_batt0 + C.uint(i)
		if C.ioctl3(fd, C.apmio_getpwstatus, &aps) < 0 {
			return "", fmt.Errorf("Failed to ioctl")
		}

		if aps.ap_batt_flag&C.apm_batt_not_present == 0 {
			icon := "    "
			if info.ai_acline == 1 {
				icon = "   "
			}
			return fmt.Sprintf("%s%v%%", icon, info.ai_batt_life), nil
		}
	}
	return "", nil
}

func audioApplet() (string, error) {
	return "", nil
}

func brightnessApplet() (string, error) {
	return "", nil
}
