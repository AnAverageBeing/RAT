package antianalysis

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// IsDebuggerPresent checks if a debugger is attached to the current process
// on both Windows and Linux.
func IsDebuggerPresent() bool {
	if runtime.GOOS == "windows" {
		isDebuggerPresent := false

		// Using the Windows API call IsDebuggerPresent
		kernel32, err := syscall.LoadLibrary("kernel32.dll")
		if err != nil {
			return false
		}
		defer syscall.FreeLibrary(kernel32)

		isDebuggerPresentProc, err := syscall.GetProcAddress(kernel32, "IsDebuggerPresent")
		if err != nil {
			return false
		}

		isDebuggerPresentResult, _, _ := syscall.SyscallN(uintptr(isDebuggerPresentProc), 0, 0, 0, 0)
		isDebuggerPresent = isDebuggerPresentResult != 0

		// Additional checks for common debuggers on Windows
		debuggers := []string{
			"ollydbg.exe",
			"idaq.exe",
			"idaq64.exe",
			"ida.exe",
			"ImmunityDebugger.exe",
			"Wireshark.exe",
		}

		for _, debugger := range debuggers {
			if processExists(debugger) {
				fmt.Printf("Debugger '%s' detected!\n", debugger)
				return true
			}
		}

		return isDebuggerPresent
	} else if runtime.GOOS == "linux" {
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
				fmt.Printf("Debugger '%s' detected!\n", debugger)
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
				fmt.Println("Debugger detected!")
				return true
			}
		}
	}

	return false
}
