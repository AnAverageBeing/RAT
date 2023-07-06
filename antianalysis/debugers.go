package antianalysis

import (
	"syscall"
)

// IsDebuggerPresent checks if a debugger is attached to the current process on Windows.
func IsDebuggerPresent() bool {
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
	}

	for _, debugger := range debuggers {
		if processExists(debugger) {
			return true
		}
	}

	return isDebuggerPresent
}
