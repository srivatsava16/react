package main

import (
	"fmt"
	"os"
)

func w() {
	v, err := os.Create("d.text")
	if err != nil {
		fmt.Println(err)
	}
	c := []string{"jwfnjkwnf,wefhinwioej,wefuh"}
	fmt.Fprintln(v, c)
	v.Close()

}

func main() {
	w()
	a, err := os.Create("f.text")
	if err != nil {
		fmt.Println(err)
	}
	_, err = a.WriteString("king")
	if err != nil {
		fmt.Println(err)
		a.Close()
	}

	a.Close()
	if err != nil {
		fmt.Println(err)
	}
}
