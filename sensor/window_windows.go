package sensor

// vim:ts=4

import (
	"log"
	"syscall"
	"unsafe"
)

type HWND uintptr

var user32 *syscall.DLL
var getForegroundWindow_W32, getWindowText_W32 *syscall.Proc
var testHandle HWND

func init() {
	user32 = syscall.MustLoadDLL("user32.dll")
	getForegroundWindow_W32 = user32.MustFindProc("GetForegroundWindow")
	getWindowText_W32 = user32.MustFindProc("GetWindowTextA")
}

func getForegroundWindow() HWND {
	if testHandle != 0 {
		return testHandle
	}
	windowHandle, _, err := getForegroundWindow_W32.Call()
	if windowHandle == 0 {
		panic("error getting foreground window handle: " + err.Error())
	}
	log.Printf("windowHandle is %v\n", windowHandle)
	return HWND(windowHandle)
}

func WindowTitle() string {
	// or you can handle the errors in the above if you want to provide some alternative
	windowHandle := getForegroundWindow()
	var buffer [256]byte
	windowTitleLen, _, _ := getWindowText_W32.Call(uintptr(windowHandle), uintptr(unsafe.Pointer(&buffer)), uintptr(256))
	if windowTitleLen == 0 {
		return ""
	}
	return string(buffer[:windowTitleLen])
}
