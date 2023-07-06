package antianalysis

// IsNetworkAnalysisRunning checks if a network analysis tool is running on Windows.
// It returns true if a network analysis tool is detected, otherwise false.
func IsNetworkAnalysisRunning() bool {
	networkAnalysisTools := []string{
		"Wireshark.exe",
		"tcpdump.exe",
		"tshark.exe",
		"WinDump.exe",
	}

	// Check if any network analysis tool is running
	for _, tool := range networkAnalysisTools {
		if processExists(tool) {
			return true
		}
	}

	return false
}
