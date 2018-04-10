package main

import "fmt"
import "../07_http_server/server"
type RootHandler struct {
}

func (h *RootHandler) Handle(request *Request) {
	fmt.Println("handle the request")
}
func main() {
	var server = Server{":8080", make(map[string]Handler)}
	server.BindHandler("/", &RootHandler{})
	server.ListenAndServe()
}
