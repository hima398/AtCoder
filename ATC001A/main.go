package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var dx = []int{1, -1, 0, 0}

var b [500][500]bool

func MakeBoolField(h, w int) [][]bool {
	f := make([][]bool, h)
	for i := 0; i < h; i++ {
		f[i] = make([]bool, w)
	}
	return f
}

func MakeByteField(h, w int) [][]byte {
	f := make([][]byte, h)
	for i := 0; i < h; i++ {
		f[i] = make([]byte, w)
	}
	return f
}

func PrintByteField(h, w int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Printf("%t ", b[i][j])
		}
		fmt.Println("")
	}
}

func SolveStack() {

}

func SolveRecursive(f []string, x, y int) {
	if x < 0 || x >= len(f) || y < 0 || y >= len(f[0]) {
		return
	}
	if f[x][y] == '#' {
		return
	}
	if b[x][y] {
		return
	}
	b[x][y] = true
	SolveRecursive(f, x+1, y)
	SolveRecursive(f, x-1, y)
	SolveRecursive(f, x, y+1)
	SolveRecursive(f, x, y-1)

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, _ := nextInt(), nextInt()
	f := make([]string, h)
	var sx, sy, gx, gy int
	for i := 0; i < h; i++ {
		f[i] = nextString()
		for j := range f[i] {
			if f[i][j] == 's' {
				sx, sy = i, j
			}
			if f[i][j] == 'g' {
				gx, gy = i, j
			}
		}
	}
	SolveRecursive(f, sx, sy)
	//PrintByteField(h, w)
	if b[gx][gy] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
