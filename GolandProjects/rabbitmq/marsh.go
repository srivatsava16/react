package main

import (
	"encoding/json"
	"fmt"
)

type emp struct {
	Name  string
	Place string
}

func main() {
	e := emp{"vat", "hyd"}
	marsh, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marsh))
}
