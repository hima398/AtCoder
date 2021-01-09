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

	s, p := nextInt(), nextInt()

	for n := 1; n*n <= p; n++ {
		if p%n != 0 {
			continue
		}
		m := p / n
		if n+m == s {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
	return
}
