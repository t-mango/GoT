package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("")
	// ssh.Start()
	service := ":8888"
	IPAddr, err := net.ResolveIPAddr("ip4", service)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	IPconn, err := net.ListenIP("ip:tcp", IPAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		// listen2Client function will listen forever until terminate
		// up to you to define how listen2Client works...
		byt := make([]byte, 1024)
		index, _ := IPconn.Read(byt)

		fmt.Println(string(byt[0:index]))
	}
}
