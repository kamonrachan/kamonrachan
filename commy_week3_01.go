package main

import "fmt"

func main() {
	n, e := fmt.Print("kamonrachan", "tokoksung", 013, 123, "\n")
	fmt.Print("number of bytes written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")
}
