package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	favourite [2]string
}

func main() {
	//var abc person
	//var a [2]string
	//a[0] = "something"
	//a[1] = "i dont know"
	//abc = person{"j", "hj", a}
	////abc.favourite= []{"sd","sd"}
	//
	//fmt.Println("valueis", abc)
	var a = make(map[int]int)
	a[1] = 1
	a[2] = 2
	fmt.Println(a)
	for k, v := range a {
		fmt.Println(k, v)
	}

}
