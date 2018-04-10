package main

import (
	"fmt"
	"bufio"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	check(err)

	defer li.Close()
	for {
		conn, err := li.Accept()
		check(err)

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		t := scanner.Text()

		if t == "" {
			fmt.Println("Receive terminate end")
			break
		}

		fmt.Printf("Get (%s)\n", t)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


