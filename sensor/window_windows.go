package sensor

// vim:ts=4

import (
	"errors"
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

type HWND uintptr
type DWORD uint32
type TickCount DWORD
type LastInputInfo struct {
	cbSize uint32
	dwTime TickCount
}

type MouseMovePoint struct {
	X, Y        int
	Time        DWORD // should be TickCount?
	dwExtraInfo *uint32
}

var (
	user32                   = syscall.MustLoadDLL("user32.dll")
	getForegroundWindow_W32  = user32.MustFindProc("GetForegroundWindow")
	getWindowText_W32        = user32.MustFindProc("GetWindowTextW")
	getLastInputInfo_W32     = user32.MustFindProc("GetLastInputInfo")
	getMouseMovePointsEx_W32 = user32.MustFindProc("GetMouseMovePointsEx")
)
var testHandle HWND

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
	var buffer [256]uint16
	windowTitleLen, _, _ := getWindowText_W32.Call(uintptr(windowHandle),
		uintptr(unsafe.Pointer(&buffer)), uintptr(256))
	if windowTitleLen == 0 {
		return ""
	}
	return syscall.UTF16ToString(buffer[:windowTitleLen])
}

func GetLastInputInfo() (TickCount, error) {
	lastInputInfo := LastInputInfo{}
	lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
	rc, _, _ := getLastInputInfo_W32.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if int(rc) == 0 {
		return 0, errors.New("No time returned")
	}
	return lastInputInfo.dwTime, nil
}

var dummyMMP = MouseMovePoint{}

const (
	GMMP_USE_DISPLAY_POINTS = iota + 1
	GMMP_USE_HIGH_RESOLUTION_POINTS
)

func GetMouseMovePointsEx() (*MouseMovePoint, error) {
	fakeMMP := MouseMovePoint{2201, 0, 0, nil}
	resultMMP := new([32]MouseMovePoint)
	rcPtr, _, _ := getMouseMovePointsEx_W32.Call(uintptr(unsafe.Sizeof(dummyMMP)),
		uintptr(unsafe.Pointer(&fakeMMP)),
		uintptr(unsafe.Pointer(&resultMMP[0])),
		32, GMMP_USE_DISPLAY_POINTS)
	rc := int32(rcPtr)
	if rc == -1 || rc != 1 {
		return nil, errors.New(fmt.Sprintf("No position returned (%d != 1)", rc))
	}
	return &resultMMP[0], nil
}
