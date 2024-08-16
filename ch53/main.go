package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	processName := "Xshell.exe"

	fmt.Println("Waiting for process", processName, "to start...")

	var processHandle windows.Handle
	go func() {
		for {
			processId, err := getProcessIdByName(processName)
			fmt.Println(processId)
			if err == nil {
				processHandle, err = windows.OpenProcess(0x000F0000|0x00100000|0xFFF, false, 7336) //uint32(processId)

				if err == nil && processHandle != 0 {
					fmt.Println("Successfully opened handle for process", processName)
					break
				}
			}

			time.Sleep(time.Second)
		}
	}()
	for {
		for i := 0; i < 0xffffffff; i += 2 {
			fmt.Printf("0x%x\n", i)
			var buffer [1024]byte

			// 监控 ReadProcessMemory()
			bytesRead := uint32(0)
			gt := uintptr(unsafe.Pointer(&bytesRead))
			success := windows.ReadProcessMemory(processHandle,
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

// 根据进程名获取进程 ID
func getProcessIdByName(processName string) (int, error) {
	snapshotHandle, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return 0, err
	}
	defer windows.CloseHandle(snapshotHandle)
	var processEntry windows.ProcessEntry32
	processEntry.Size = uint32(unsafe.Sizeof(processEntry))
	err = windows.Process32First(snapshotHandle, &processEntry)
	if err != nil {
		return 0, err
	}
	for {
		if syscall.UTF16ToString(processEntry.ExeFile[:]) == processName {
			return int(processEntry.ProcessID), nil
		}
		err = windows.Process32Next(snapshotHandle, &processEntry)
		if err != nil {
			break
		}
	}
	return 0, fmt.Errorf("process not found")
}
