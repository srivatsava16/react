//In this file the execution of all APIs of our project will be done in this file

package main

import (
	dao "awesomeProject1/db"
	entities "awesomeProject1/functions"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.Databaseconn()
	a := gin.Default()
	a.POST("/createnewtask", entities.Createtask)
	a.POST("/createnewemp", entities.Createemp)
	a.PUT("/updateprojectid/:projectid", entities.Update)
	a.GET("/task4", entities.Monthfilter)
	a.GET("/task5", entities.Task5)
	a.Run()
}
