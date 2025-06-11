package main

import "fmt"

func Add10(addend *int) {
	*addend = *addend + 10
}

func main() {
	c := new(int)
	*c = 5
	Add10(c)
	fmt.Println(*c)

	s := &[]int{2, 3}
	DoubleSlice(s)
	fmt.Println(s)
}
