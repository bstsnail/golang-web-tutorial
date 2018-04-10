package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"io"
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

	var i = 1

	for scanner.Scan() {

		t := scanner.Text()

		if t == "" {
			fmt.Println("Recevie terminate end")
			break
		}

		if i == 1 {
			fmt.Printf("Get the request (%s)\n", t)
			words := strings.Split(t, " ")

			if len(words) == 1 {
				fmt.Fprintln(conn, "World")
				return
			}
			method := words[0]
			path := words[1]

			if path == "/" {
				index(conn)
			} else if path == "/dog" {
				if method == "GET" {
					dog(conn)
				} else if method == "POST" {
					postDog(conn)
				} else {
					fmt.Printf("Unknown method(%s)\n", method)
				}
			} else {
				fmt.Printf("Unknown path(%s)\n", path)
			}
		}
		i++
	}
}

func index(c net.Conn) {
	fmt.Fprintln(c, "Hello world")
}

func dog(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html>
		<body>
		<form action="/dog" method="POST">
		<input name="fname" type="text" placeholder="first name">
		<input name="lname" type="text" placeholder="last name">
		<input type="submit" value="SEND TO SERVER">
		</form>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func postDog(c net.Conn) {
	fmt.Fprintln(c, "Post dog")
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
