package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func (p Point) X() int { return p.x }
func (p Point) Y() int { return p.y }

func main() {
	// s := "hello"
	// s[0] = 'H'

	//可修改
	a := [3]int{1, 2, 3}
	a[0] = 4
	a[0]++
	b := [3]int{4, 5, 6}
	a = b
	a[0] = 7
	b[0] = 8

	fmt.Println(a)
	fmt.Println(b)
	
	p := Point{1, 2}
	p.x = 3
}
