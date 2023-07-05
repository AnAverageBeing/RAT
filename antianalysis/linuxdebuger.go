//go:build linux
// +build linux

package antianalysis

import (
	"os"
	"os/exec"
	"strings"
)

// IsDebuggerPresent checks if a debugger is attached to the current process on Linux.
func IsDebuggerPresent() bool {
	// Check for common debuggers on Linux
	debuggers := []string{
		"gdb",
		"lldb",
		"strace",
		"ltrace",
		"processhacker",
	}

	for _, debugger := range debuggers {
		if _, err := exec.LookPath(debugger); err == nil {
			return true
		}
	}

	// Check if a debugger is attached
	if _, err := os.Stat("/proc/self/status"); err == nil {
		file, err := os.Open("/proc/self/status")
		if err != nil {
			return false
		}
		defer file.Close()

		buf := make([]byte, 2048)
		count, _ := file.Read(buf)

		if strings.Contains(string(buf[:count]), "TracerPid:\t0") {
			return true
		}
	}

	return false
}
