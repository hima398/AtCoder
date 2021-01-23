package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

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

type Pos struct {
	X int
	Y int
}

type Stack struct {
	p []Pos
}

func NewStack() *Stack {
	return new(Stack)
}

func (this *Stack) Push(p Pos) {
	this.p = append(this.p, p)
}

func (this *Stack) Pop() *Pos {
	n := len(this.p)
	if n == 0 {
		return nil
	}
	p := this.p[n-1]
	this.p = this.p[:n-1]
	return &p
}

func SolveStack(f []string, x, y int) {

	stack := NewStack()
	p := Pos{x, y}
	stack.Push(p)
	b[x][y] = true
	for {
		p := stack.Pop()
		b[p.X][p.Y] = true
		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]
			if nx >= 0 && nx < len(f) && ny >= 0 && ny < len(f[0]) && !b[nx][ny] && f[nx][ny] != '#' {
				np := Pos{nx, ny}
				stack.Push(np)
			}
		}
		if len(stack.p) == 0 {
			break
		}
	}
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
	for i := 0; i < 4; i++ {
		SolveRecursive(f, x+dx[i], y+dy[i])
	}
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
	//SolveRecursive(f, sx, sy)
	SolveStack(f, sx, sy)
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
