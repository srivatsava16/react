package main

import (
	"bytes"
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	buf := bytes.NewBuffer(make([]byte, 0, 1e5))

	err := client.SetImageFromBytes("image_crop_mad-0.1004806682471.png")
	if err != nil {
		fmt.Println("parsing mage", err)
	}

	text, _ := client.HOCRText()
	fmt.Println(text)
}
