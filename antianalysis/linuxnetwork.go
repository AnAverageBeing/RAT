//go:build linux
// +build linux

package antianalysis

import (
	"os/exec"
)

// IsNetworkAnalysisRunning checks if a network analysis tool is running on Linux.
// It returns true if a network analysis tool is detected, otherwise false.
func IsNetworkAnalysisRunning() bool {
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
			return true
		}
	}

	return false
}
