//In this file the execution of all APIs of our project will be done in this file

package main

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/db/functions"
)

func main() {
	dao.Databaseconn()
	a := gin.Default()
	a.POST("/createnewtask", tasks.Createtask)
	a.POST("/createnewemp", tasks.Createemp)
	a.PUT("/updateprojectid/:projectid", tasks.Update)
	a.GET("/task4", tasks.Monthfilter)
	a.GET("/task5", tasks.Task5)

	a.Run()
}
