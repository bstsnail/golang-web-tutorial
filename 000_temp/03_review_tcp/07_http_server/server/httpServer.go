package server

import (
	"net"
	"container/list"
	"io"
	"fmt"
	"bufio"
	"strings"
)

type Request struct {
	Method string
	Headers *list.List
	Body string
	Uri string
	Query *list.List
	conn net.Conn
}

func (request *Request) send(response string) {
	io.WriteString(request.conn, response)
}

type Handler interface {
	Handle(request *Request)
}


type Server struct {
	Address string
	Handlers map[string]Handler
}

func (server *Server) BindHandler(uri string, handler Handler) {
	server.Handlers[uri] = handler
}

func (server *Server) handle(request *Request)  {
	h, exist := server.Handlers[request.Uri]
	if !exist {
		return
	}
	h.Handle(request)
}



func (server *Server) ListenAndServe() {

	listener, err := net.Listen("tcp", server.Address)

	check(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		check(err)
		fmt.Printf("Recevie a request from %s\n", conn.RemoteAddr().String())
		go serve(server, conn)
	}

}

func serve(server *Server, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	start := true

	var method string
	var uri string
	headers := list.New()
	query := list.New()

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		if start {
			words := strings.Split(t, " ")
			method = words[0]
			path := words[1]
			index := strings.Index(path, "?")
			if index == -1 {
				uri = path
			} else {
				uri = path[0:index]
				queries := path[index:]

				for index = strings.Index(queries, "&"); index != -1; {
					query.PushBack(queries[0:index])
					queries = queries[index:]
				}
				query.PushBack(queries)
			}
			start = false
		} else {
			words := strings.Split(t, ":")
			m := make(map[string]string)
			m[words[0]] = strings.Trim(words[1], " ")
			headers.PushBack(m)
		}
	}

	request := &Request{
		method,
		headers,
		"",
		uri,
		query,
		conn,
	}
	server.handle(request)
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
