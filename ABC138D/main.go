package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Stack struct {
	v []int
}

func NewStack() *Stack {
	return new(Stack)
}

func (this *Stack) Push(v int) {
	this.v = append(this.v, v)
}

func (this *Stack) Pop() int {
	n := len(this.v)
	if n == 0 {
		return 0
	}
	v := this.v[n-1]
	this.v = this.v[:n-1]
	return v
}

func (this *Stack) Size() int {
	return len(this.v)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	r := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		r[a] = append(r[a], b)
		r[b] = append(r[b], a)
	}
	p := make([]int, q)
	x := make(map[int]int)

	for i := 0; i < q; i++ {
		p[i] = nextInt()
		x[p[i]] += nextInt()
	}
	stack := NewStack()
	cnt := make([]int, n+1)
	visited := make([]bool, n+1)

	stack.Push(1)
	for {
		par := stack.Pop()
		if visited[par] {
			continue
		}
		cnt[par] += x[par]
		visited[par] = true
		for _, ch := range r[par] {
			x[ch] += x[par]
			stack.Push(ch)
		}
		if stack.Size() == 0 {
			break
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", cnt[i])
	}
	fmt.Println("")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
