package main

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type describe interface {
	describe() string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	fmt.Printf("co={num : %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num: ", co.base.num)

	fmt.Println("describe: ", co.describe())

	var d describe = co

	fmt.Println("describe: ", d.describe())
}
