package main

import (
	"fmt"
	"net"
	"bufio"
	"io"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	check(err)

	defer li.Close()

	for {
		conn, err := li.Accept()

		check(err)

		fmt.Printf("Receive connection from %s\n", conn.RemoteAddr().String())

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			t := scanner.Text()
			fmt.Printf("Get %s\n", t)

			if t == "" {
				break
			}
		}

		io.WriteString(conn, "hello")
		conn.Close()
	}


}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
