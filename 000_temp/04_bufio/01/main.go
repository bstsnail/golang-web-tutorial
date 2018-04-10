package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf("Get text (%s)\n", t)
	}
}
