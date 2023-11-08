package funct

import (
	"awesomeProject4/struct"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var (
	hitCount int
	hitQueue = make([]models.Manualresponse, 8)
)

func ToQueue(hitname string, input models.Student, response models.Response) error {
	var hitData models.Manualresponse
	hitData.Hit = hitname
	hitData.Input = input
	hitData.Manual = response
	if hitCount > 7 {
		return errors.New("hits out of bounds")
	}
	if hitCount == 7 {
		hitQueue[hitCount] = hitData
		hitCount++
		//fmt.Println(hitQueue[hitCount])
		err := ToFile(hitQueue)
		return err
	} else {
		hitQueue[hitCount] = hitData
		//fmt.Println(hitQueue[hitCount])
		hitCount++
	}
	return nil
}
func ToFile(hitQueue []models.Manualresponse) error {

	fh, err := os.OpenFile("api.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	log.Println("File opened")
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			log.Println(err)
		}
	}(fh)

	jsonObj, err := json.MarshalIndent(&hitQueue, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(fh, string(jsonObj))
	if err != nil {
		return err
	}
	log.Println(string(rune(hitCount)) + " hits written to file")
	return nil
}

func Stud(c *gin.Context) {
	var stu models.Student
	var response models.Response

	stu = models.Student{
		Id:        c.PostForm("Id"),
		Firstname: c.PostForm("firstname"),
		Lastname:  c.PostForm("lastname"),
		Username:  c.PostForm("username"),
		Password:  c.PostForm("password"),
		Address:   c.PostForm("address"),
		Place:     c.PostForm("place"),
	}

	if stu.Id == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "Id cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Id empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if stu.Firstname == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "Firstname cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Firstname empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if stu.Lastname == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "Lastname cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Lastname empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if stu.Password == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "password cannot be empty",
			Data:   nil,
		}

		err := ToQueue("password empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if stu.Username == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "username cannot be empty",
			Data:   nil,
		}

		err := ToQueue("username empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if stu.Address == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "address cannot be empty",
			Data:   nil,
		}

		err := ToQueue("address empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if stu.Place == "" {

		response = models.Response{
			Status: "Error",
			Code:   400,
			Error:  "place cannot be empty",
			Data:   nil,
		}

		err := ToQueue("place empty", stu, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = models.Response{
		Status: "Success",
		Code:   200,
		Error:  "",
		Data:   "success",
	}

	err := ToQueue("Success hit", stu, response)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, response)
	return
}
