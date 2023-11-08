package main

import "fmt"

type AK struct {
	vat string
}

func main() {
	var a = make(map[string][]AK)
	a["s"] = []AK{AK{"S"}, AK{"S"}}
	s, _ := fmt.Println(a)
	fmt.Println(s)
}
