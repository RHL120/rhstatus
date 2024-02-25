#ifdef  __FreeBSD__
#include <fcntl.h>
#include <unistd.h>
#include <machine/apm_bios.h>
#include <sys/soundcard.h>
#include <sys/backlight.h>
#include "c.h"

int get_battery(battery_t *stat) {
	int i;
	bool ret = -1;
	struct apm_info info = {0};
	struct apm_pwstatus pw_stat = {0};
	int fd = open("/dev/apm", O_RDONLY);
	if (fd < 0) {
		goto ret;
	}
	if (ioctl(fd, APMIO_GETINFO, &info) < 0) {
		goto close_fd;
	}
	for (i = 0; i < info.ai_batteries; i++) {
		pw_stat.ap_device = PMDV_BATT0 + i;
		if (ioctl(fd, APMIO_GETPWSTATUS, &pw_stat) < 0) {
			goto close_fd;
		}
		if (!(pw_stat.ap_batt_flag & APM_BATT_NOT_PRESENT)) {
			stat->charging = info.ai_acline == 1;
			stat->charge = info.ai_batt_life;
			ret = 1;
			goto close_fd;
		}

	}
	ret = 0;
close_fd: close(fd);
ret: return ret;
}

int get_sound() {
	int ret = -1;
	int dev_info = 0;
	int fd = open("/dev/mixer", O_RDONLY);
	if (fd < 0) {
		goto ret;
	}
	if (ioctl(fd, MIXER_READ(0), &dev_info) < 0) {
		goto close_fd;
	}
	ret = dev_info & 0x7f;
close_fd: close(fd);
ret: return ret;
}

int get_brightness() {
	struct backlight_props props = {0};
	int ret = -1;
	int fd = open("/dev/backlight/backlight0", O_RDONLY);
	if (fd < 0) {
		goto ret;
	}
	if (ioctl(fd, BACKLIGHTGETSTATUS, &props) < 0) {
		goto close_fd;
	}
	ret = props.brightness;
close_fd: close(fd);
ret: return ret;
}
#endif
