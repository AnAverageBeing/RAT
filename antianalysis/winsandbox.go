//go:build windows
// +build windows

package antianalysis

import (
	"syscall"
)

func SandBoxDetected() bool {
	dlls := []string{
		"SbieDll",
		"SxIn",
		"Sf2",
		"snxhk",
		"cmdvrt32",
	}

	for _, dll := range dlls {
		handle, err := syscall.LoadLibrary(dll + ".dll")
		if err == nil && handle != 0 {
			defer syscall.FreeLibrary(handle)
			return true
		}
	}

	return false
}
