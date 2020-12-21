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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

var Pyramid [100][100]int

type Coordinate struct {
	X int
	Y int
	H int
}

func Check(x, y int, cs []Coordinate) (int, bool) {
	ch := cs[0].H + Abs(cs[0].X-x) + Abs(cs[0].Y-y)
	for _, c := range cs {
		h := Max(ch-Abs(x-c.X)-Abs(y-c.Y), 0)
		if h != c.H {
			return 0, false
		}
	}
	return ch, true
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	c := make([]Coordinate, n+1)
	for i := 1; i <= n; i++ {
		c[i].X = nextInt()
		c[i].Y = nextInt()
		c[i].H = nextInt()
		if c[i].H > 0 {
			c[0] = c[i]
		}
	}

	for cx := 0; cx <= 100; cx++ {
		for cy := 0; cy <= 100; cy++ {
			h, ok := Check(cx, cy, c)
			if ok {
				fmt.Printf("%d %d %d\n", cx, cy, h)
				return
			}
		}
	}
}
