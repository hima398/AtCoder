package main

import "fmt"

var nn, mm []int

func PrintScore(n int) {
	println(nn[n-1])
}

func SolveProbrem(n, m int) {
	nn[n-1]++
	mm[m-1]--
}

func main() {
	var n, m, q int

	fmt.Scan(&n, &m, &q)

	nn = make([]int, n, n)
	mm = make([]int, m, m)
	for i := 0; i < m; i++ {
		mm[i] = n
	}
	fmt.Println(nn)
	fmt.Println(mm)
	var s, nnn, mmm int
	for i := 0; i < q; i++ {
		fmt.Scan(&s, &nnn, &mmm)
		if s == 1 {
			PrintScore(nnn)
		}
		if s == 2 {
			SolveProbrem(nnn, mmm)
		}
	}
}
