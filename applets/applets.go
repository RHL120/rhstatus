package applets

type Applet struct {
	update_interval uint
	enabled         bool
	function        func(...interface{}) (string, error)
}

var Applets map[string]Applet = map[string]Applet{
	"date": {update_interval: 60, enabled: true, function: date_applet},
	"time": {update_interval: 60, enabled: true, function: time_applet},
}

func init() {

}
