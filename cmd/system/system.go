package main

import (
	"fmt"
	"github.com/TheDen/system"
	"log"
)

func main() {
	filesystem := "/"
	inodes, err := system.Inodes(filesystem)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(inodes)
	data, err := system.Uptime()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
