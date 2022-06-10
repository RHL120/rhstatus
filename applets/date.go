package applets

import "time"

const (
	datefmt = " Monday January 2"
	timefmt = " 15:04"
)

func dateApplet(...interface{}) (string, error) {
	current := time.Now()
	return current.Format(datefmt), nil
}
func timeApplet(...interface{}) (string, error) {
	current := time.Now()
	return current.Format(timefmt), nil
}
