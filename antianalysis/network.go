package antianalysis

import (
	"fmt"
	"os/exec"
	"runtime"
)

// IsNetworkAnalysisRunning checks if a network analysis tool is running.
// It returns true if a network analysis tool is detected, otherwise false.
func IsNetworkAnalysisRunning() bool {
	if runtime.GOOS == "windows" {
		networkAnalysisTools := []string{
			"Wireshark.exe",
			"tcpdump.exe",
			"tshark.exe",
			"WinDump.exe",
		}

		// Check if any network analysis tool is running
		for _, tool := range networkAnalysisTools {
			if processExists(tool) {
				fmt.Printf("Network analysis tool '%s' detected!\n", tool)
				return true
			}
		}
	} else if runtime.GOOS == "linux" {
		// Check for common network analysis tools on Linux
		networkAnalysisTools := []string{
			"tcpdump",
			"wireshark",
			"tshark",
			"netstat",
			"netmon",
		}

		// Check if any network analysis tool is in the system's PATH
		for _, tool := range networkAnalysisTools {
			if _, err := exec.LookPath(tool); err == nil {
				fmt.Printf("Network analysis tool '%s' detected!\n", tool)
				return true
			}
		}
	}

	return false
}
