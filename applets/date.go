package applets

import "time"

const (
	datefmt = "Monday January 2"
	timefmt = "15:04"
)

func date_applet(...interface{}) (string, error) {
	current := time.Now()
	return current.Format(datefmt), nil
}
func time_applet(...interface{}) (string, error) {
	current := time.Now()
	return current.Format(timefmt), nil
}
