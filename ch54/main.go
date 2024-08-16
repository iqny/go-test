package main

import (
	"fmt"
	"unsafe"
)

import (
	windows "golang.org/x/sys/windows"
)

var (
	data   [2]byte
	length uint32
)

func main() {
	handle, _ := windows.OpenProcess(0x0010, false, 7336) // 0x0010 PROCESS_VM_READ, PID 6100
	procReadProcessMemory := windows.MustLoadDLL("kernel32.dll").MustFindProc("ReadProcessMemory")

	for i := 0; i < 0xffffffff; i += 2 {
		fmt.Printf("0x%x\n", i)

		// BOOL ReadProcessMemory(HANDLE hProcess, LPCVOID lpBaseAddress, LPVOID lpBuffer, DWORD nSize, LPDWORD lpNumberOfBytesRead)
		ret, _, e := procReadProcessMemory.Call(uintptr(handle), uintptr(i),
			uintptr(unsafe.Pointer(&data[0])),
			2, uintptr(unsafe.Pointer(&length))) // read 2 bytes
		if ret == 0 {
			fmt.Println("  Error:", e)
		} else {
			fmt.Println("  Length:", length)
			fmt.Println("  Data:", data)
		}
	}

	windows.CloseHandle(handle)
}
