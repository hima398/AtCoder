package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Road struct {
	To int
	D  int
}

func Push(s []int, v int) []int {
	s = append(s, v)
	return s
}

func Pop(s []int) (int, []int) {
	n := len(s)
	if n == 0 {
		return 0, s
	}
	v := s[n-1]
	ns := s[:n-1]
	return v, ns
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	// 入力
	n := nextInt()
	r := make(map[int][]Road)
	for i := 0; i < n-1; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		atb := Road{b, c}
		bta := Road{a, c}
		r[a] = append(r[a], atb)
		r[b] = append(r[b], bta)
	}
	q, k := nextInt(), nextInt()
	x := make([]int, q)
	y := make([]int, q)
	for i := 0; i < q; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}

	//
	var stack []int
	d := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		d[i] = -1
	}
	d[k] = 0

	stack = Push(stack, k)
	for {
		var i int
		i, stack = Pop(stack)

		for _, v := range r[i] {
			if d[v.To] == -1 {
				d[v.To] = d[i] + v.D
				stack = Push(stack, v.To)
			}
		}
		if len(stack) == 0 {
			break
		}
	}

	for i := 0; i < q; i++ {
		fmt.Println(d[x[i]] + d[y[i]])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
