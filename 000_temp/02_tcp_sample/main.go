package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", "baidu.com:80")

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Status=" + status)
}
