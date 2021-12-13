//go:build darwin || freebsd || netbsd || openbsd
// +build darwin freebsd netbsd openbsd

// Package system returns system information
package system

// Get the system uptime returned in seconds
func Uptime() (int64, error) {
	return 1, nil
}
