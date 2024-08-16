package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32                   = syscall.NewLazyDLL("user32.dll")
	procFindWindowW          = user32.NewProc("FindWindowW")
	procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")
	procGetWindowTextW       = user32.NewProc("GetWindowTextW")
)

func FindWindow(lpClassName, lpWindowName *uint16) (hwnd syscall.Handle) {
	ret, _, _ := procFindWindowW.Call(
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)))
	hwnd = syscall.Handle(ret)
	return
}

func GetWindowTextLength(hwnd syscall.Handle) int32 {
	ret, _, _ := procGetWindowTextLengthW.Call(uintptr(hwnd))
	return int32(ret)
}

func GetWindowText(hwnd syscall.Handle, lpString *uint16, nMaxCount int32) int32 {
	ret, _, _ := procGetWindowTextW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(nMaxCount))
	return int32(ret)
}

func main() {
	// 假设我们正在寻找标题为"计算器"的窗口
	className, _ := syscall.UTF16PtrFromString("CalcFrame")
	windowName, _ := syscall.UTF16PtrFromString("计算器")

	hwnd := FindWindow(className, windowName)
	if hwnd == 0 {
		fmt.Println("未找到窗口")
		return
	}

	length := GetWindowTextLength(hwnd)
	buf := make([]uint16, length+1)
	GetWindowText(hwnd, &buf[0], int32(len(buf)))

	windowTitle := syscall.UTF16ToString(buf)
	fmt.Printf("窗口标题: %s\n", windowTitle)
}
