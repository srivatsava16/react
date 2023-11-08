package main

import (
	"awesomeProject4/funct"
	"github.com/gin-gonic/gin"
)

func main() {
	a := gin.Default()
	a.POST("/newstuu", funct.Stud)

	a.Run()

}
