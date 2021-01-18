package main

import (
	"fmt"
	"time"
)

// SayStatement ...
func SayStatement(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func main() {
	go SayStatement("World")
	fmt.Println("Hello")
	time.Sleep(3 * time.Second)
	fmt.Println("End of Program")
}
