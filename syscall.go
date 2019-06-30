// Package system returns system information
package system

import (
	"syscall"
)

// Get the free file nodes in the filesystem
func Inodes(fs string) (int64, error) {
	var data syscall.Statfs_t
	err := syscall.Statfs(fs, &data)
	if err != nil {
		return 0, err
	}
	return int64(data.Ffree), nil
}

// Get the system uptime returned in seconds
func Uptime() (int64, error) {
	var data syscall.Sysinfo_t
	err := syscall.Sysinfo(&data)
	if err != nil {
		return data.Uptime, err
	}
	return data.Uptime, nil
}

// Get user identity
func Userid() int {
	return syscall.Getuid()
}

// Get group identity
func Groupid() int {
	return syscall.Getgid()
}

// Kill a process
func Kill(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}
