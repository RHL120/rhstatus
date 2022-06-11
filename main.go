package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include <stdlib.h>
// #define DefaultScreen(dpy) 	(((_XPrivDisplay)(dpy))->default_screen)
// #define RootWindow(dpy, scr) (ScreenOfDisplay(dpy,scr)->root)
// unsigned long rw (Display *dpy) {
//return RootWindow(dpy, DefaultScreen(dpy));
//}
//void render(Display *dpy, Window root, const char *status) {
//XStoreName(dpy, root, status);
//XFlush(dpy);
//}
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/RHL120/rhstatus/applets"
)

func main() {
	dpy := C.XOpenDisplay(nil)
	root := C.rw(dpy)
	defer C.XCloseDisplay(dpy)
	for {
		var goName string
		for name, i := range applets.Applets {
			ret, err := i.Function()
			fmt.Println(name)
			if err != nil {
				fmt.Printf("Failed to run applet %s because %v\n", name, err)
				continue
			}
			goName = fmt.Sprintf("%s  |  %s", goName, ret)
		}
		name := C.CString(goName)
		C.render(dpy, root, name)
		C.free(unsafe.Pointer(name))
		time.Sleep(5 * time.Second)
	}

}
