package main

import (
	"fmt"
	"os"
	"syscall"
)

func Get(fs string) (int, error) {
	var data syscall.Statfs_t
	err := syscall.Statfs(fs, &data)
	if err != nil {
		return 0, err
	}
	return int(float64(data.Files-data.Ffree) / float64(data.Files) * 100.0), nil
}

func main() {
	args := os.Args
	filesystem := "/"
	if len(args) > 1 {
		filesystem = os.Args[1]
	}
	percentageused, err := Get(filesystem)
	if err != nil {
		panic(err)
	}
	fmt.Println(percentageused)
}
