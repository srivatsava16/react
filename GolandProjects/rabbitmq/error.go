package main

import "fmt"

type err struct {
	message string
}

func (e err) Error() string {
	return e.message
}
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &err{
			"b cant be 0",
		}
	} else if a == 0 {
		return 0, &err{
			"0",
		}
	} else {
		return a + b, nil
	}
}
func main() {
	x, y := 0, 880
	a, err := div(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(a)
	}
}
