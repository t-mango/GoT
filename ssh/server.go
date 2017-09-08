package ssh

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func Start() {

	//var hostKey ssh.PublicKey

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("w1ngti5zhu"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", "115.159.125.233:22", config)

	if err != nil {
		log.Fatal("链接错误", err)
	}
	defer conn.Close()

	l, err := conn.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("链接错误", err)
	}
	defer l.Close()

	for {
		//net.TCPListener()
		conn, err := l.Accept()

	}

	// http.Serve(l, http.HandlerFunc(
	// 	func(resp http.ResponseWriter, req *http.Request) {
	// 		fmt.Fprintf(resp, "mango")
	// 	}))
}
