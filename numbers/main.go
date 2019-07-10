package main

import "fmt"

func main() {
	n := numbers{
		first:  4,
		second: 3,
	}
	p := n.multiply()
	fmt.Println(p)
}

type numbers struct {
	first  int
	second int
}

func (n *numbers) multiply() (p int) {
	p = n.first * n.second
	return
}
