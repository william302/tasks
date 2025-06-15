package main

import (
	"fmt"
)

func Podd() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func Peven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	scheduler([]func(){
		randomtask,
		randomtask,
		randomtask,
	})
}
