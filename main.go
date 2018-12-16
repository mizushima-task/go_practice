package main

import (
	"fmt"
)

type Hoge struct {
	x, y, z int
}

func main() {
	v := Hoge{1, 2, 3}
	p := v
	fmt.Println(p)
}
