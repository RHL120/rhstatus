package X

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
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

func OpenDisplay() *C.Display {
	return C.XOpenDisplay(nil)
}

func CloseDisplay(dpy *C.Display) {
	C.XCloseDisplay(dpy)
}

func UpdateStatus(dpy interface{}, status string) {
	rw := C.rw(dpy.(*C.Display))
	cstatus := C.CString(status)
	defer C.free(unsafe.Pointer(cstatus))
	C.XStoreName(dpy.(*C.Display), rw, cstatus)
	C.XFlush(dpy.(*C.Display))
}
