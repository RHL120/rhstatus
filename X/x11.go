package X

// #cgo linux LDFLAGS: -lX11
// #cgo openbsd CFLAGS: -I/usr/X11R6/include
// #cgo openbsd LDFLAGS: -L/usr/X11R6/lib -lX11
// #ifdef __linux__
//   #include <X11/Xlib.h>
// #elif defined(__OpenBSD__)
//   #include <X11/Xlib.h>
//   #include <X11/Xutil.h>
// #endif
// #include <stdlib.h>
// #define DefaultScreen(dpy) 	(((_XPrivDisplay)(dpy))->default_screen)
// #define RootWindow(dpy, scr) (ScreenOfDisplay(dpy,scr)->root)
// unsigned long rw (Display *dpy) {
//return RootWindow(dpy, DefaultScreen(dpy));
//}
//void render(Display *dpy, Window root, const char *status) {
//	XStoreName(dpy, root, status);
//	XFlush(dpy);
//}
import "C"
import "unsafe"

var dpy = C.XOpenDisplay(nil)
var rw = C.rw(dpy)

//open a display
func OpenDisplay() *C.Display {
	return C.XOpenDisplay(nil)
}

//close the display
func CloseDisplay() {
	C.XCloseDisplay(dpy)
}

//clear the status bar and put status onto it
func UpdateStatus(status string) {
	rw := C.rw(dpy)
	cstatus := C.CString(status)
	defer C.free(unsafe.Pointer(cstatus))
	C.XStoreName(dpy, rw, cstatus)
	C.XFlush(dpy)
}
