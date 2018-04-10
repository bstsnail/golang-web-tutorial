package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	check(err)

	r4 := bufio.NewReader(f)

	b4, err := r4.Peek(4)
	check(err)
	fmt.Printf("Get the str(%s)", string(b4))

}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
