//All the API in our project is written in this file

package entities

import (
	"awesomeProject1/db"
	"awesomeProject1/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Createtask(c *gin.Context) {
	var newtask models.Task

	Id, _ := strconv.Atoi(c.PostForm("id"))
	AssignedEmployeeId, _ := strconv.Atoi(c.PostForm("AssignedEmployeeId"))
	Resolved, _ := strconv.Atoi(c.PostForm("Resolved"))
	ProjectId, _ := strconv.Atoi(c.PostForm("ProjectId"))
	newtask = models.Task{
		Id:                 Id,
		Task:               c.PostForm("Task"),
		Type:               c.PostForm("Type"),
		CreatedAt:          c.PostForm("Createdat"),
		AssignedEmployeeId: AssignedEmployeeId,
		Resolved:           Resolved,
		ProjectId:          ProjectId,
	}
	if err := c.Bind(&newtask); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error-Bind": err.Error()})
		return
	}

	err := dao.DB.Create(&newtask).Error
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{

		"data":   newtask,
		"status": "Successfully created",
		"code":   "200",
		"error":  "",
	})
}

func Createemp(c *gin.Context) {
	var newemp models.Employee
	Id, _ := strconv.Atoi(c.PostForm("Id"))
	Onboardingcompletion, _ := strconv.Atoi(c.PostForm("Onboardingcompletion"))
	newemp = models.Employee{
		Id:                   Id,
		FirstName:            c.PostForm("Firstname"),
		LastName:             c.PostForm("Lastname"),
		FullName:             c.PostForm("Fullname"),
		Jobtitle:             c.PostForm("Jobtitle"),
		Jobtype:              c.PostForm("Jobtype"),
		Phone:                c.PostForm("Phone"),
		Email:                c.PostForm("Email"),
		Image:                c.PostForm("Image"),
		Country:              c.PostForm("Country"),
		City:                 c.PostForm("City"),
		Onboardingcompletion: Onboardingcompletion,
	}
	if err := c.Bind(&newemp); err != nil {

		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error-Bind": err.Error()})
		return
	}

	err := dao.DB.Create(&newemp).Error
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{

		"data":   newemp,
		"status": "Successfully created",
		"code":   "200",
		"error":  "",
	})
}

func Update(c *gin.Context) {
	var update models.Task
	id := c.Param("projectid")
	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id should not be empty"})
		return
	}
	if err := c.ShouldBindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"if found": err.Error()})
		return
	}
	err := dao.DB.Where("id=?", id).Updates(&update).Error
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "not valid"})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{

		"data":   update,
		"status": "Successfully created",
		"code":   "200",
		"error":  "",
	})
}

func Monthfilter(c *gin.Context) {
	var monthfilter models.Task
	curdate := c.Request.FormValue("month")
	pdate, err := time.Parse("2006-01-02", curdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	twomonthsback := pdate.AddDate(0, -2, 0)
	err1 := dao.DB.Model(&models.Task{}).Where("created_at BETWEEN ? AND ?", twomonthsback, pdate).Preload("Employees").Preload("Projects").Find(&monthfilter).Error
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"data":   monthfilter,
		"status": "Successfully ",
		"code":   "200",
		"error":  "",
	})
}

func Task5(c *gin.Context) {
	var a []models.Task

	from := c.Request.FormValue("from")
	fromm, err := time.Parse("2006-01-02", from)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	to := c.Request.FormValue("to")
	too, err := time.Parse("2006-01-02", to)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if from == "" || to == "" {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	if err := dao.DB.Model(&models.Task{}).Where("created_at  BETWEEN ? AND ?", fromm, too).Preload("Projects", "Employees").Find(&a).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}
	c.JSON(http.StatusOK, gin.H{

		"data":   a,
		"status": "Successfully created",
		"code":   "200",
		"error":  "",
	})
}
