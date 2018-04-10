package main

import (
	"fmt"
	"github.com/bstsnail/golang-web-tutorial/000_temp/03_review_tcp/07_http_server/server"
)

type RootHandler struct {
}

func (h *RootHandler) Handle(request *server.Request) {
	fmt.Println("handle the request")
}
func main() {
	s := &server.Server{
		Address: ":8080",
		Handlers: make(map[string]server.Handler),
	}
	s.BindHandler("/", &RootHandler{})
	s.ListenAndServe()
}
