// Package system returns system information for linux
package main

import (
	"fmt"
	"log"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/TheDen/system"
	"golang.org/x/sys/unix"
)

func main() {
	os := runtime.GOOS
	fmt.Println(os)

	filesystem := "/"
	inodes, err := system.Inodes(filesystem)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(inodes)

	data, _ := system.Uptime()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	out, err := unix.SysctlRaw("kern.boottime")
	if err != nil {
		fmt.Println(time.Duration(0), err)
	}
	var timeval syscall.Timeval
	if len(out) != int(unsafe.Sizeof(timeval)) {
		fmt.Println(time.Duration(0), fmt.Errorf("unexpected output of sysctl kern.boottime: %v (len: %d)", out, len(out)))
	}
	timeval = *(*syscall.Timeval)(unsafe.Pointer(&out[0]))
	sec, nsec := timeval.Unix()
	fmt.Println(time.Now().Sub(time.Unix(sec, nsec)))
}
