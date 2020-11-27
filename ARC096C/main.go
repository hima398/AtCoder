package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func Compare(A, B *Purchase) (*Purchase, *Purchase) {
	if A.quantity < B.quantity {
		return A, B
	}
	return B, A
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Purchase struct {
	price    int
	quantity int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c, x, y := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	A := &Purchase{a, x}
	B := &Purchase{b, y}

	total := a*x + b*y
	greater := c * Max(x, y) * 2
	Less, Greater := Compare(A, B)
	less := c*Less.quantity*2 + Greater.price*(Greater.quantity-Less.quantity)

	prices := []int{total, greater, less}
	sort.Ints(prices)

	fmt.Println(prices[0])
}
