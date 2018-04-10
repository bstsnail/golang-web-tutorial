package main

import (
	"net"
	"io"
	"fmt"
)

func main() {

	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err.Error())
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Recevie a connection ")
		io.WriteString(conn, "DO IT WORK?")
		conn.Close()
	}
}
