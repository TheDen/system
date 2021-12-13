package system

import "syscall"

// Get the system uptime returned in seconds
func Uptime_linux() (int64, error) {
	var data syscall.Sysinfo_t
	err := syscall.Sysinfo(&data)
	if err != nil {
		return data.Uptime, err
	}
	return data.Uptime, nil
}
