package antianalysis

import (
	"os"
	"strings"
)

// IsVm checks if the application is running in a virtual machine environment
func IsVm() bool {
	// Check for common virtualization tools
	virtualizationTools := []string{
		"/dev/kvm",          // Kernel-based Virtual Machine (KVM)
		"/dev/vhost-net",    // Linux Kernel Virtual Machine (KVM) with vhost-net module
		"/dev/vboxdrv",      // Oracle VirtualBox
		"/dev/vmmon",        // VMware
		"/dev/uhid",         // User-level HID (uhid) interface
		"/dev/xen",          // Xen Hypervisor
		"/dev/virtio-ports", // Virtio serial ports
	}

	for _, tool := range virtualizationTools {
		if _, err := os.Stat(tool); err == nil {
			return true
		}
	}

	// Check for some general VM indicators
	indicators := []string{
		"/sys/devices/virtual/misc/meminfo", // Memory info file for virtual devices
		"/proc/fb",                          // Framebuffer devices
		"/proc/cpuinfo",                     // CPU information
		"/proc/scsi/scsi",                   // SCSI devices
		"/sys/class/dmi/id/product_name",    // Product name from DMI table
		"/sys/class/dmi/id/sys_vendor",      // System vendor from DMI table
	}

	for _, indicator := range indicators {
		if _, err := os.Stat(indicator); err == nil {
			return true
		}
	}

	// Check for hypervisor in CPU flags
	cpuInfo, err := os.Open("/proc/cpuinfo")
	if err == nil {
		defer cpuInfo.Close()

		cpuFlags := make([]byte, 4096)
		_, err := cpuInfo.Read(cpuFlags)
		if err == nil && strings.Contains(string(cpuFlags), "hypervisor") {
			return true
		}
	}

	return false
}
