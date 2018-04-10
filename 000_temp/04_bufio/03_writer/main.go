package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	f, err := os.Create("file.txt")
	check(err)

	defer f.Close()

	n1, err := f.Write([]byte("hello go"))
	check(err)
	fmt.Printf("Write %d bytes to file\n", n1)

	f.Sync()

	w := bufio.NewWriter(f)
	n2, err := w.WriteString("Hello go by buffer")
	check(err)

	fmt.Printf("Write %d bytes\n", n2)
	w.Flush()

}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
