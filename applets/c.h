#ifndef APPS_C_C
#define  APPS_C_C
#include <stdbool.h>
typedef struct {
	bool charging;
	int charge;
} battery_t;
/// Fills *stat*, returns:
/// -1: if an error occured
/// 0 : if the battery is absent
/// 1 : otherwise
int get_battery(battery_t *stat);
/// Returns the volume or -1 if an error occured
int get_sound();
/// Returns the brightness or -1 if an error occured
int get_brightness();
#endif
