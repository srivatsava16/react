package main

import (
	funct "awesomeProject1/func"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	a := gin.Default()
	a.POST("/newemp", funct.Createemp)
	err := a.Run()
	if err != nil {

		log.Println("unable to run")
	}
}
