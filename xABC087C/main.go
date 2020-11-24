package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	if n == 1 {
		a1, a2 := nextInt(), nextInt()
		fmt.Println(a1 + a2)
		return
	}

	// n >= 2
	a1 := make([]int, n)
	a2 := make([]int, n)
	top := make([]int, n+1)
	bottom := make([]int, n+1)
	for i := 0; i < n; i++ {
		a1[i] = nextInt()
		top[i+1] = a1[i] + top[i]
	}
	for i := 0; i < n; i++ {
		a2[i] = nextInt()
		bottom[i+1] = a2[i] + bottom[i]
	}
	var i int
	for i = 0; i < n; i++ {
		if top[n]-top[i+1]+a2[n-1] < bottom[n]-bottom[i] {
			break
		}
	}
	if i == n {
		fmt.Println(top[i] + a2[n-1])
		return
	}
	fmt.Println(top[i+1] + (bottom[n] - bottom[i]))
	return
}
