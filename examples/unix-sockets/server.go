package main

import (
	"fmt"
	"net"
)

func handledConn(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		fmt.Println("Reading..")
		num, err := conn.Read(buf)

		if err != nil {
			fmt.Println("NO DATA")
			return
		}

		fmt.Println("reading ", num, ": ", string(buf))
	}
}

func main() {
	fmt.Println("Initialization...")

	lis, _ := net.Listen("unix", "server.sock")

	for {
		conn, _ := lis.Accept()

		go handledConn(conn)
	}
}
