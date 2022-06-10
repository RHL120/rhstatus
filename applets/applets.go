package applets

type Applet struct {
	name            string
	icon            string
	update_interval string
	function        func(...interface{}) (string, error)
}
