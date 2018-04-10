package main

import (
	"fmt"
	"io"
	"net"
	"io/ioutil"
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

		fmt.Println("Recevive a connection from " + conn.RemoteAddr().String())

		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Get " + string(bs))

		io.WriteString(conn, "hello")

		conn.Close()
	}
}
