package main

import "fmt"

func main() {
	S := "U"
	N := 100000
	for i := 2; i < N-1; i++ {
		S += "U"
	}
	S += "D"
	fmt.Printf("%s", S)
}
