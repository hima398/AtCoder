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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	p := make([]int, n+1)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = nextInt()
		d[i] = p[i] - i
	}
	//fmt.Println(p)
	//fmt.Println(d)

	ans := 0
	for i := 1; i < n; i++ {
		if d[i] == 0 {
			p[i], p[i+1] = p[i+1], p[i]
			d[i] = p[i] - i
			d[i+1] = p[i+1] - (i + 1)
			ans++
		}
	}
	if d[n] == 0 {
		p[n-1], p[n] = p[n], p[n-1]
		d[n-1] = p[n-1] - (n - 1)
		d[n] = p[n] - n
		ans++
	}
	//fmt.Println(p)
	//fmt.Println(d)
	fmt.Println(ans)
}
