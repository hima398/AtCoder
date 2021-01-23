package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	X int
	Y int
	D int
}

type Queue struct {
	ps []Pos
}

func NewQueue() *Queue {
	return new(Queue)
}

func (this *Queue) Push(p Pos) {
	this.ps = append(this.ps, p)
}

func (this *Queue) Pop() *Pos {
	if len(this.ps) == 0 {
		return nil
	}
	p := this.ps[0]
	this.ps = this.ps[1:]
	return &p
}

func (this *Queue) Find(x, y int) bool {
	for _, v := range this.ps {
		if x == v.X && y == v.Y {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	r, c := nextInt(), nextInt()
	sy, sx := nextInt()-1, nextInt()-1
	gy, gx := nextInt()-1, nextInt()-1
	f := make([]string, r)
	d := make([][]int, r)
	for i := 0; i < r; i++ {
		f[i] = nextString()
		d[i] = make([]int, c)
		for j := 0; j < c; j++ {
			d[i][j] = -1
		}
	}
	queue := NewQueue()
	sp := Pos{sx, sy, 0}
	d[sp.Y][sp.X] = sp.D
	queue.Push(sp)
	for {
		p := queue.Pop()
		d[p.Y][p.X] = p.D
		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]
			if nx >= 0 && nx < c && ny >= 0 && ny < r && f[ny][nx] != '#' && d[ny][nx] == -1 && !queue.Find(nx, ny) {
				np := Pos{nx, ny, p.D + 1}
				queue.Push(np)
			}
		}

		if queue.Size() == 0 {
			break
		}
	}
	/*
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				fmt.Printf("%d ", d[i][j])
			}
			fmt.Println("")
		}
	*/
	fmt.Println(d[gy][gx])
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
