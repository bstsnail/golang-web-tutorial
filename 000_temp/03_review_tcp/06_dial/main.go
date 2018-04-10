package main

import (
	"fmt"
	"net"
	"bufio"
	"io"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	check(err)
	defer c.Close()

	io.WriteString(c, "Hello\n")

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
