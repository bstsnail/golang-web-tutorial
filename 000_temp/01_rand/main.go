package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(2)
	fmt.Println(x)
}
