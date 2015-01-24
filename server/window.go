package main

// vim:ts=4

import (
	"fmt"
	"syscall"
	"unsafe"
)

func WindowTitle() string {
	user32 := syscall.MustLoadDLL("user32.dll")                       // or NewLazyDLL() to defer loading
	getForegroundWindow := user32.MustFindProc("GetForegroundWindow") // or NewProc() if you used NewLazyDLL()
	getWindowText := user32.MustFindProc("GetWindowTextA")            // or NewProc() if you used NewLazyDLL()
	// or you can handle the errors in the above if you want to provide some alternative
	r1, _, err := getForegroundWindow.Call()
	// err will always be non-nil; you need to check r1 (the return value)
	if r1 == 0 { // in this case
		panic("error getting foreground window handle: " + err.Error())
	}
	var buffer [256]byte
	r2, _, err := getWindowText.Call(r1, uintptr(unsafe.Pointer(&buffer)), uintptr(256))
	if r2 == 0 {
		return ""
	}
	// fmt.Printf("%v\n", buffer)
	return string(buffer[:r2])
}

func main() {
	fmt.Printf("Hello, world.\n")
}
