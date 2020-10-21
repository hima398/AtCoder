package main

import "fmt"

func GeometricProgression(a, r, n int) int {
	// 10 ^ 9
	val := 1000000000
	g := a
	for i := 0; i < n-1; i++ {
		g *= r
		if g > val {
			return -1
		}
	}
	return g
}

func main() {
	var a, r, n int
	fmt.Scan(&a, &r, &n)
	g := GeometricProgression(a, r, n)
	if g == -1 {
		fmt.Println("large")
	} else {
		fmt.Println(g)
	}
}
