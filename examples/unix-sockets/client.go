package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Initialization...")

	conn, _ := net.Dial("unix", "server.sock")

	for {
		conn.Write([]byte("hola mundo"))

		time.Sleep(time.Second)

	}

}
