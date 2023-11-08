package main

import (
	"errors"
	"fmt"
)

func main() {
	newerr := errors.New("SET")
	a, b := 1, 0
	if b == 0 {
		fmt.Println(newerr, b)
		//panic("set")

	} else {
		fmt.Println(a / b)

	}
	fmt.Println(10 / 5)

}
