package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32             = syscall.MustLoadDLL("user32.dll")
	procEnumWindows    = user32.MustFindProc("EnumWindows")
	procGetWindowTextW = user32.MustFindProc("GetWindowTextW")
)

func EnumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FindWindow(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := GetWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		if syscall.UTF16ToString(b) == title {
			// note the window
			hwnd = h
			return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	EnumWindows(cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("No window with title '%s' found", title)
	}
	return hwnd, nil
}

func main() {
	const title = "微信支付"
	h, err := FindWindow(title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found '%s' window: handle=0x%x\n", title, h)
	//var processHandle windows.Handle

	for {
		for i := 0; i < 0xffffffff; i += 4 {
			fmt.Printf("0x%x\n", i)
			var buffer [1024]byte

			// 监控 ReadProcessMemory()
			bytesRead := uint32(0)
			gt := uintptr(unsafe.Pointer(&bytesRead))
			success := windows.ReadProcessMemory(windows.Handle(h),
				uintptr(i),
				&buffer[0],
				uintptr(len(buffer)),
				&gt)

			if success == nil && gt > 0 {
				fmt.Printf("[Intercepted] ReadProcessMemory: %s\n", string(buffer[:gt]))
			} else {
				fmt.Println("read err:", success)
			}

			// 监控 WriteProcessMemory()
			/*bytesWritten := uint32(0)
			bytesWrittenGt := uintptr(unsafe.Pointer(&bytesWritten))
			success = windows.WriteProcessMemory(processHandle,
				uintptr(unsafe.Pointer(&buffer)),
				&buffer[0],
				uintptr(len(buffer)),
				&bytesWrittenGt)

			if success == nil && bytesWrittenGt > 0 {
				fmt.Printf("[Intercepted] WriteProcessMemory: %s\n", string(buffer[:bytesWrittenGt]))
			} else {
				fmt.Println("write err:", success)
			}*/
		}
		time.Sleep(time.Second)

	}
}
