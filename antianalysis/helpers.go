package antianalysis

import (
	"fmt"
	"os/exec"
	"strings"
)

// Helper function to check if a process is running
func processExists(processName string) bool {
	cmd := exec.Command("tasklist", "/NH", "/FI", fmt.Sprintf("IMAGENAME eq %s", processName))
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.Contains(string(output), processName)
}
