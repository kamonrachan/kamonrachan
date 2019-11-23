package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	commystart1 := time.Now()
	var a string
	for i := 0; i < 50000; i++ {
		a += "a"
	}
	fmt.Println("a", time.Since(commystart1))

	commystart2 := time.Now()
	var b strings.Builder
	for i := 0; i < 50000; i++ {
		b.WriteString("b")
	}
	fmt.Println("b", time.Since(commystart2))
}
